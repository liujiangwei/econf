<template>
  <div>
    <a-form layout="inline" @submit="onSearch" :form="form">
      <a-form-item label="前缀匹配">
        <a-switch
          checkedChildren="是"
          unCheckedChildren="否"
          defaultChecked
          v-decorator="['prefix', {rule:[], initialValue: true}]"
        />
      </a-form-item>
      <a-form-item>
        <a-input placeholder="input key" v-decorator="['key', {rule:[], initialValue: '/'}]" />
      </a-form-item>
      <a-form-item>
        <a-button type="primary" html-type="submit">查询</a-button>
      </a-form-item>
    </a-form>
    <a-table :dataSource="kvs" :columns="columns" >
      <a slot="name" slot-scope="text" href="javascript:;">{{key}}</a>
      <span slot="customTitle">
        <a-icon type="smile-o" />Name
      </span>
      <span slot="tags" slot-scope="tags">
        <a-tag v-for="tag in tags" color="blue" :key="tag">{{value}}</a-tag>
      </span>
      <span slot="tags" slot-scope="tags">
        <a-tag v-for="tag in tags" color="blue" :key="tag">{{version}}</a-tag>
      </span>
    </a-table>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "Etcd",
  data() {
    return {
      form: this.$form.createForm(this),
      columns: [
        {
          title: "键名",
          dataIndex: "key",
          key: "key"
        },
        {
          title: "键值",
          dataIndex: "value",
          key: "value"
        },
        {
          title: "版本",
          dataIndex: "version",
          key: "version"
        },
        {
          title: "租约",
          dataIndex: "lease",
          key: "lease"
        },
        {
          title: "租约时间",
          dataIndex: "granted_ttl",
          key: "granted_ttl"
        },
        {
          title: "剩余时间",
          dataIndex: "ttl",
          key: "ttl"
        }
      ],
      kvs: []
    };
  },
  beforeCreate() {},
  mounted: function() {},
  methods: {
    onSearch(e) {
      e.preventDefault();
      //console.log(this.form.getFieldsValue())

      var data = {
        key: this.form.getFieldValue("key"),
        prefix: this.form.getFieldValue("prefix") ? "1" : "0"
      };
      axios
        .get("http://localhost:8081/etcd/search", { params: data })
        .then(response => {
          this.kvs = response.data;
          //console.log(response)
        })
        .catch(error => {
          console.log(error);
        });
    }
  }
};
</script>

<style scoped>
</style>