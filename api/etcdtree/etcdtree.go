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

	resp, err := client.Get(context.Background(), key, options...)
	if err != nil{
		return err
	}

	var data = make([]map[string]interface{}, 0)

	for _, kv := range resp.Kvs{
		data = append(data, map[string]interface{}{
			"key":string(kv.Key),
			"value": string(kv.Value),
			"version": kv.Version,
		})
	}

	return c.JSON(http.StatusOK, data)
}

type tree struct {
	Item []item
}

type item struct {
	Key string
	Value string
	Version int64
}