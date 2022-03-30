/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	infrastructurev1beta1 "github.com/outscale-vbr/cluster-api-provider-outscale.git/api/v1beta1"
	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/record"
	"time"
        "os"
	//      "k8s.io/apimachinery/pkg/runtime"
	"github.com/outscale-vbr/cluster-api-provider-outscale.git/cloud/scope"
        "github.com/outscale-vbr/cluster-api-provider-outscale.git/util/reconciler"
	"sigs.k8s.io/cluster-api/util"
	"sigs.k8s.io/cluster-api/util/annotations"
	ctrl "sigs.k8s.io/controller-runtime"
        "github.com/outscale-vbr/cluster-api-provider-outscale.git/cloud/services/service" 
        "github.com/outscale-vbr/cluster-api-provider-outscale.git/cloud/services/net"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// OscClusterReconciler reconciles a OscCluster object
type OscClusterReconciler struct {
	client.Client
	Recorder         record.EventRecorder
	ReconcileTimeout time.Duration
}

//+kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=oscclusters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=oscclusters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=oscclusters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the OscCluster object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *OscClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (_ ctrl.Result, reterr error) {
	_ = log.FromContext(ctx)
	ctx, cancel := context.WithTimeout(ctx, reconciler.DefaultedLoopTimeout(r.ReconcileTimeout))
	defer cancel()
	log := ctrl.LoggerFrom(ctx)
	oscCluster := &infrastructurev1beta1.OscCluster{}

	log.Info("Please WAIT !!!!")

	if err := r.Get(ctx, req.NamespacedName, oscCluster); err != nil {
		if apierrors.IsNotFound(err) {
    			log.Info("object was not found")
			return ctrl.Result{}, nil
		}
          	return ctrl.Result{}, err
        }
	log.Info("Still WAIT !!!!")
        log.Info("Create info", "env", os.Environ())

	cluster, err := util.GetOwnerCluster(ctx, r.Client, oscCluster.ObjectMeta)
	if err != nil {
		return reconcile.Result{}, err
	}
	if cluster == nil {
		log.Info("Cluster Controller has not yet set OwnerRef")
		return reconcile.Result{}, nil
	}

	// Return early if the object or Cluster is paused.
	if annotations.IsPaused(cluster, oscCluster) {
		log.Info("oscCluster or linked Cluster is marked as paused. Won't reconcile")
		return ctrl.Result{}, nil
	}

	// Create the cluster scope.
	clusterScope, err := scope.NewClusterScope(scope.ClusterScopeParams{
		Client:     r.Client,
		Logger:     log,
		Cluster:    cluster,
		OscCluster: oscCluster,
	})
	if err != nil {
		return reconcile.Result{}, errors.Errorf("failed to create scope: %+v", err)
	}
	defer func() {
		if err := clusterScope.Close(); err != nil && reterr == nil {
			reterr = err
		}
	}()
	osccluster := clusterScope.OscCluster
	if !osccluster.DeletionTimestamp.IsZero() {
		return r.reconcileDelete(ctx, clusterScope)
	}
        loadBalancerSpec := clusterScope.LoadBalancer()
        loadBalancerSpec.SetDefaultValue()
        log.Info("Create loadBalancer", "loadBalancerName", loadBalancerSpec.LoadBalancerName, "SubregionName", loadBalancerSpec.SubregionName)
	return r.reconcile(ctx, clusterScope)
}


func (r *OscClusterReconciler) reconcile(ctx context.Context, clusterScope *scope.ClusterScope) (reconcile.Result, error) {
    clusterScope.Info("Reconcile OscCluster")
    osccluster := clusterScope.OscCluster
    servicesvc := service.NewService(ctx, clusterScope)
    clusterScope.Info("Get Service", "service", servicesvc)

    clusterScope.Info("Create Loadbalancer")
    loadBalancerSpec := clusterScope.LoadBalancer()
    loadBalancerSpec.SetDefaultValue()
    loadbalancer, err := servicesvc.GetLoadBalancer(loadBalancerSpec)
    if err != nil {
        return reconcile.Result{}, err
    }
    if loadbalancer == nil {
    	_, err := servicesvc.CreateLoadBalancer(loadBalancerSpec)
	if err != nil {
            return reconcile.Result{}, errors.Wrapf(err, "Can not create load balancer for Osccluster %s/%s", osccluster.Namespace, osccluster.Name)
    	}
        _, err = servicesvc.ConfigureHealthCheck(loadBalancerSpec)
        if err != nil {
            return reconcile.Result{}, errors.Wrapf(err, "Can not configure healthcheck for Osccluster %s/%s", osccluster.Namespace, osccluster.Name)
        } 
    }
    netsvc := net.NewService(ctx, clusterScope)
    clusterScope.Info("Get net", "net", netsvc)

    clusterScope.Info("Create Net")
    netSpec := clusterScope.Net()
    netSpec.SetDefaultValue()
    netRef := clusterScope.NetRef()
    netName := "cluster-api-net-" + clusterScope.UID()
    if len(netRef.ResourceMap) == 0 {
        netRef.ResourceMap = make(map[string]string)
    }
    var netIds = []string{netRef.ResourceMap[netName]}
    net, err := netsvc.GetNet(netIds)
    clusterScope.Info("### len net ###", "net", len(netRef.ResourceMap))
    clusterScope.Info("### Get net ###", "net", net)
    clusterScope.Info("### Get netIds ###", "net", netIds)
    if err != nil {
        return reconcile.Result{}, err
    }
    if net == nil {
        clusterScope.Info("### Empty Net ###")
        netRef.ResourceMap[netName] = "init"
        clusterScope.Info("### content net ###", "net", netRef.ResourceMap)
        net, err = netsvc.CreateNet(netSpec, netName)
        if err != nil {
            return reconcile.Result{}, errors.Wrapf(err, "Can not create load balancer for Osccluster %s/%s", osccluster.Namespace, osccluster.Name)
        }
        netRef.ResourceMap[netName] = *net.NetId
        clusterScope.Info("### content updatee net ###", "net", netRef.ResourceMap)

    }
    netRef.ResourceMap[netName] = *net.NetId
    clusterScope.Info("Info net", "net", net)

    clusterScope.Info("Create Subnet")
    subnetSpec := clusterScope.Subnet()
    subnetSpec.SetDefaultValue()
    subnetRef := clusterScope.SubnetRef()   
    subnetName := "cluster-api-subnet-" + clusterScope.UID()
    if len(subnetRef.ResourceMap) == 0 {
        subnetRef.ResourceMap = make(map[string]string)
    }
    var subnetIds = []string{subnetRef.ResourceMap[subnetName]}
    subnet, err := netsvc.GetSubnet(subnetIds)
    clusterScope.Info("### len subnet ###", "subnet", len(subnetRef.ResourceMap))
    clusterScope.Info("### Get subnet ###", "subnet", subnet)
    clusterScope.Info("### Get subnetIds ###", "subnet", subnetIds)
    if err != nil {
        return reconcile.Result{}, err
    }
    if subnet == nil {
        clusterScope.Info("### Empty Subnet ###") 
        subnetRef.ResourceMap[subnetName] = "init"
        clusterScope.Info("### content subnet ###", "subnet", subnetRef.ResourceMap)
        subnet, err = netsvc.CreateSubnet(subnetSpec, netRef.ResourceMap[netName], subnetName)
        if err != nil {
            return reconcile.Result{}, errors.Wrapf(err, "Can not create subnet for Osccluster %s/%s", osccluster.Namespace, osccluster.Name)
        }
        subnetRef.ResourceMap[subnetName] = *subnet.SubnetId
        clusterScope.Info("### content update subnet ###", "subnet", subnetRef.ResourceMap)
    }

    clusterScope.Info("Create InternetGateway")
    internetServiceSpec := clusterScope.InternetService()
    internetServiceRef := clusterScope.InternetServiceRef()
    internetServiceName := "cluster-api-internetservice-" + clusterScope.UID()
    if len(internetServiceRef.ResourceMap) == 0 {
        internetServiceRef.ResourceMap = make(map[string]string)
    }
    var internetServiceIds = []string{internetServiceRef.ResourceMap[internetServiceName]}
    internetService, err := netsvc.GetInternetService(internetServiceIds)
    clusterScope.Info("### len internetService ###", "internetservice", len(internetServiceRef.ResourceMap))
    clusterScope.Info("### Get internetService ###", "internetservice",  internetService)
    clusterScope.Info("### Get internetServiceIds ###", "internetservice",  internetServiceIds)
    if err != nil {
        return reconcile.Result{}, err
    }
    if internetService == nil {
        clusterScope.Info("### Empty internetService ###")
        internetServiceRef.ResourceMap[internetServiceName] = "init"
        clusterScope.Info("### content internetService ###", "internetservice", internetServiceRef.ResourceMap)
        internetService, err = netsvc.CreateInternetService(internetServiceSpec, internetServiceName)
        if err != nil {
            return reconcile.Result{}, errors.Wrapf(err, "Can not create internetservice for Osccluster %s/%s", osccluster.Namespace, osccluster.Name)
        }
        err = netsvc.LinkInternetService(*internetService.InternetServiceId, netRef.ResourceMap[netName])
        if err != nil {
            return reconcile.Result{}, errors.Wrapf(err, "Can not link internetService with net for Osccluster %s/%s", osccluster.Namespace, osccluster.Name)
        }
        internetServiceRef.ResourceMap[internetServiceName] = *internetService.InternetServiceId
        clusterScope.Info("### content update internetService ###", "internetservice", internetServiceRef.ResourceMap)

    }

    clusterScope.Info("Create RouteTable")
    routeTablesSpec := clusterScope.RouteTables()
    routeTablesRef := clusterScope.RouteTablesRef()
    var routeTableIds []string
    for _, routeTableSpec := range *routeTablesSpec {
        routeTableName := routeTableSpec.Name + clusterScope.UID()
        routeTableIds = []string{routeTablesRef.ResourceMap[routeTableName]}    
        if len(routeTablesRef.ResourceMap) == 0 {
            routeTablesRef.ResourceMap = make(map[string]string)
        }
        routeTable, err := netsvc.GetRouteTable(routeTableIds)
        clusterScope.Info("### Get routeTable ###", "routeTable", routeTable)
        clusterScope.Info("### Get routeTableIds ###", "routeTable",  routeTableIds)
        if err != nil {
            return reconcile.Result{}, err
        }
        if routeTable == nil {
            clusterScope.Info("### Empty routeTable ###")
            routeTablesRef.ResourceMap[routeTableName] = "init"
            clusterScope.Info("### content routeTable ###", "routeTable", routeTablesRef.ResourceMap)
            routeTable, err = netsvc.CreateRouteTable(&routeTableSpec, netRef.ResourceMap[netName], routeTableName)
            if err != nil {
                return reconcile.Result{}, errors.Wrapf(err, "Can not create internetservice for Osccluster %s/%s", osccluster.Namespace, osccluster.Name)
            }
            routeTablesRef.ResourceMap[routeTableName] = *routeTable.RouteTableId
            clusterScope.Info("### content update routeTable ###", "routeTable", routeTablesRef.ResourceMap)
        }
    }

    controllerutil.AddFinalizer(osccluster, "oscclusters.infrastructure.cluster.x-k8s.io")
    clusterScope.Info("Set OscCluster status to ready")
    clusterScope.SetReady()
    return reconcile.Result{}, nil
}

func (r *OscClusterReconciler) reconcileDelete(ctx context.Context, clusterScope *scope.ClusterScope) (reconcile.Result, error) {
    clusterScope.Info("Reconcile OscCluster")
    osccluster := clusterScope.OscCluster
    servicesvc := service.NewService(ctx, clusterScope)
    clusterScope.Info("Get Service", "service", servicesvc)
    netRef := clusterScope.NetRef()

    clusterScope.Info("Delete LoadBalancer")
    loadBalancerSpec := clusterScope.LoadBalancer()
    loadBalancerSpec.SetDefaultValue()
    netName := "cluster-api-net-" + clusterScope.UID()
    var netIds = []string{netRef.ResourceMap[netName]}
    loadbalancer, err := servicesvc.GetLoadBalancer(loadBalancerSpec)
    if err != nil {
        return reconcile.Result{}, err
    }
    if loadbalancer == nil {
        controllerutil.RemoveFinalizer(osccluster, "oscclusters.infrastructure.cluster.x-k8s.io")
        return reconcile.Result{}, nil
    }
    err = servicesvc.DeleteLoadBalancer(loadBalancerSpec)
    if err != nil {
        return reconcile.Result{}, errors.Wrapf(err, "Can not delete load balancer for Osccluster %s/%s", osccluster.Namespace, osccluster.Name)
    }
    netsvc := net.NewService(ctx, clusterScope)
    clusterScope.Info("Get Net", "net", netsvc)

    clusterScope.Info("Delete RouteTable")
    routeTablesSpec := clusterScope.RouteTables()
    routeTablesRef := clusterScope.RouteTablesRef()
    var routeTableIds []string
    for _, routeTableSpec := range *routeTablesSpec {
        routeTableSpec.SetDefaultValue()
        routeTableName := routeTableSpec.Name + clusterScope.UID()
        routeTableIds = []string{routeTablesRef.ResourceMap[routeTableName]}
        routetable, err := netsvc.GetRouteTable(routeTableIds)
        if err != nil {
            return reconcile.Result{}, err
        }
        if routetable == nil {
            controllerutil.RemoveFinalizer(osccluster, "oscclusters.infrastructure.cluster.x-k8s.io")
            return reconcile.Result{}, nil
        }
        err = netsvc.DeleteRouteTable(routeTablesRef.ResourceMap[routeTableName])
        if err != nil {
            return reconcile.Result{}, errors.Wrapf(err, "Can not delete internetService for Osccluster %s/%s", osccluster.Namespace, osccluster.Name)
        }
    }
    clusterScope.Info("Delete internetService")
    internetServiceRef := clusterScope.InternetServiceRef()
    internetServiceName := "cluster-api-internetservice-" + clusterScope.UID()
    var internetServiceIds = []string{internetServiceRef.ResourceMap[internetServiceName]}
    internetservice, err := netsvc.GetInternetService(internetServiceIds)
    if err != nil {
        return reconcile.Result{}, err
    }
    if internetservice == nil {
        controllerutil.RemoveFinalizer(osccluster, "oscclusters.infrastructure.cluster.x-k8s.io")
        return reconcile.Result{}, nil
    }
    err = netsvc.UnlinkInternetService(internetServiceRef.ResourceMap[internetServiceName], netRef.ResourceMap[netName])
    if err != nil {
         return reconcile.Result{}, errors.Wrapf(err, "Can not unlink internetService and net for Osccluster %s/%s", osccluster.Namespace, osccluster.Name)
    }
    err = netsvc.DeleteInternetService(internetServiceRef.ResourceMap[internetServiceName])
    if err != nil {
         return reconcile.Result{}, errors.Wrapf(err, "Can not delete internetService for Osccluster %s/%s", osccluster.Namespace, osccluster.Name)
    }


    clusterScope.Info("Delete subnet")
    subnetRef := clusterScope.SubnetRef()
    subnetName := "cluster-api-subnet-" + clusterScope.UID()
    var subnetIds = []string{subnetRef.ResourceMap[subnetName]}
    subnet, err := netsvc.GetSubnet(subnetIds)
    if err != nil {
        return reconcile.Result{}, err
    }
    if subnet == nil {
        controllerutil.RemoveFinalizer(osccluster, "oscclusters.infrastructure.cluster.x-k8s.io")
        return reconcile.Result{}, nil
    }
    err = netsvc.DeleteSubnet(subnetRef.ResourceMap[subnetName])
    if err != nil {
         return reconcile.Result{}, errors.Wrapf(err, "Can not delete subnet for Osccluster %s/%s", osccluster.Namespace, osccluster.Name)
    }

    clusterScope.Info("Delete net")
    net, err := netsvc.GetNet(netIds)
    if err != nil {
        return reconcile.Result{}, err
    }
    if net == nil {
        controllerutil.RemoveFinalizer(osccluster, "oscclusters.infrastructure.cluster.x-k8s.io")
        return reconcile.Result{}, nil
    }
    err = netsvc.DeleteNet(netRef.ResourceMap[netName])
    if err != nil {
        return reconcile.Result{}, errors.Wrapf(err, "Can not delete net for Osccluster %s/%s", osccluster.Namespace, osccluster.Name)
    }
    controllerutil.RemoveFinalizer(osccluster, "oscclusters.infrastructure.cluster.x-k8s.io")
    return reconcile.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *OscClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrastructurev1beta1.OscCluster{}).
		Complete(r)
}