package etcdtree

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.etcd.io/etcd/clientv3"
	"net/http"
)

func List(c echo.Context) error{
	client , err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
	})
	if err != nil{
		c.Logger().Info(err)
		return err
	}

	key := c.QueryParam("key")
	if key == ""{
		key = "/"
	}

	var options []clientv3.OpOption
	if c.QueryParam("prefix") == "1"{
		options = append(options, clientv3.WithPrefix())
	}

	//putResp, putErr := client.Put(context.Background(), "a", "b")
	resp, err := client.Get(context.Background(), key, options...)
	if err != nil{
		return err
	}

	var data []keyData

	for _, kv := range resp.Kvs{
		kd := keyData{
			Key: string(kv.Key),
			Value:string(kv.Value),
			Version: kv.Version,
			Lease: kv.Lease,
		}

		if kv.Lease > 0{
			lease := clientv3.NewLease(client)
			leaseResp, leaseErr := lease.TimeToLive(context.Background(), clientv3.LeaseID(kv.Lease))
			if leaseErr == nil{
				kd.GrantedTtl = leaseResp.GrantedTTL
				kd.Ttl = leaseResp.TTL
			}else{
				c.Logger().Info(err)
			}
		}

		data = append(data, kd)
	}

	return c.JSON(http.StatusOK, data)
}

type keyData struct {
	Key string `json:"key"`
	Value string	`json:"value"`
	Version int64 `json:"version"`
	Lease int64 `json:"lease"`
	Ttl int64 `json:"ttl"`
	GrantedTtl int64 `json:"granted_ttl"`
}