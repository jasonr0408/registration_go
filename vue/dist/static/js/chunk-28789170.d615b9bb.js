(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-28789170"],{"2ee3":function(t,e,o){"use strict";o.d(e,"h",function(){return r}),o.d(e,"e",function(){return s}),o.d(e,"g",function(){return a}),o.d(e,"a",function(){return n}),o.d(e,"f",function(){return l}),o.d(e,"b",function(){return u}),o.d(e,"i",function(){return c}),o.d(e,"d",function(){return m}),o.d(e,"c",function(){return d});var i=o("b775");function r(t){return Object(i["a"])({url:"/class",method:"post",data:t,baseURL:"/go-registration"})}function s(){return Object(i["a"])({url:"/class/all",method:"get",baseURL:"/go-registration"})}function a(t,e){return Object(i["a"])({url:"/student/import",method:"post",params:{classID:t},data:e,baseURL:"/go-registration"})}function n(t,e){return Object(i["a"])({url:"/student/check-in",method:"put",params:{classID:t},data:e,baseURL:"/go-registration"})}function l(t){return Object(i["a"])({url:"/class/status",method:"get",params:{classID:t},baseURL:"/go-registration"})}function u(t){return Object(i["a"])({url:"/class",method:"delete",params:{classID:t},baseURL:"/go-registration"})}function c(t,e){return Object(i["a"])({url:"/student",method:"post",params:{classID:t},data:e,baseURL:"/go-registration"})}function m(t,e){return Object(i["a"])({url:"/student",method:"put",params:{classID:t},data:e,baseURL:"/go-registration"})}function d(t,e){return Object(i["a"])({url:"/student",method:"delete",params:{classID:t,employeeID:e},baseURL:"/go-registration"})}},"4a1f":function(t,e,o){"use strict";o.r(e);var i=function(){var t=this,e=t.$createElement,o=t._self._c||e;return o("el-container",[o("el-main",{attrs:{height:"auto"}},[o("el-row",[o("el-col",{attrs:{span:24}},[o("el-button",{staticStyle:{width:"100%","margin-bottom":"10px"},attrs:{type:"primary",icon:"el-icon-full-screen"},on:{click:t.goToScan}},[t._v("報到")])],1)],1),t._v(" "),o("el-row",[o("el-col",{staticClass:"newMember",attrs:{span:24,align:"right"}},[o("el-button",{attrs:{type:"primary",icon:"el-icon-plus"},on:{click:function(e){t.addDialogVisible=!0}}},[t._v("新增學員")])],1)],1),t._v(" "),o("el-table",{staticStyle:{width:"100%"},attrs:{data:t.list}},[o("el-table-column",{attrs:{prop:"status",label:"狀態"},scopedSlots:t._u([{key:"default",fn:function(e){return[e.row.status?o("el-tag",{attrs:{type:"success"}},[t._v("報到")]):t._e()]}}])}),t._v(" "),o("el-table-column",{attrs:{prop:"name",label:"暱稱"}}),t._v(" "),o("el-table-column",{attrs:{prop:"group",label:"組別"}}),t._v(" "),o("el-table-column",{attrs:{prop:"department",label:"部門名稱"}}),t._v(" "),o("el-table-column",{attrs:{prop:"employeeID",label:"員工編號"}}),t._v(" "),o("el-table-column",{attrs:{label:"操作",align:"center",width:"150px"},scopedSlots:t._u([{key:"default",fn:function(e){return[o("el-button",{attrs:{type:"primary",icon:"el-icon-edit",circle:"",title:"編輯"},on:{click:function(o){return t.enterStatus(e.row)}}}),t._v(" "),o("el-button",{attrs:{type:"danger",icon:"el-icon-delete",circle:"",title:"刪除"},on:{click:function(o){return t.deleteStudentByApi(e.row)}}})]}}])})],1),t._v(" "),o("el-dialog",{attrs:{title:"新增學員",visible:t.addDialogVisible,width:"350px"},on:{"update:visible":function(e){t.addDialogVisible=e}}},[o("el-form",{ref:"form",attrs:{model:t.newForm,"label-width":"100px"}},[o("el-form-item",{attrs:{label:"狀態"}},[o("el-switch",{attrs:{"active-text":"報到","inactive-text":"未報到"},model:{value:t.newForm.status,callback:function(e){t.$set(t.newForm,"status",e)},expression:"newForm.status"}})],1),t._v(" "),o("el-form-item",{attrs:{label:"暱稱"}},[o("el-input",{model:{value:t.newForm.name,callback:function(e){t.$set(t.newForm,"name",e)},expression:"newForm.name"}})],1),t._v(" "),o("el-form-item",{attrs:{label:"組別"}},[o("el-input",{model:{value:t.newForm.group,callback:function(e){t.$set(t.newForm,"group",e)},expression:"newForm.group"}})],1),t._v(" "),o("el-form-item",{attrs:{label:"部門"}},[o("el-input",{model:{value:t.newForm.department,callback:function(e){t.$set(t.newForm,"department",e)},expression:"newForm.department"}})],1),t._v(" "),o("el-form-item",{attrs:{label:"員工編號"}},[o("el-input",{model:{value:t.newForm.employeeID,callback:function(e){t.$set(t.newForm,"employeeID",e)},expression:"newForm.employeeID"}})],1)],1),t._v(" "),o("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[o("el-button",{on:{click:function(e){t.addDialogVisible=!1}}},[t._v("取 消")]),t._v(" "),o("el-button",{attrs:{type:"primary"},on:{click:t.summitAddForm}},[t._v("送 出")])],1)],1),t._v(" "),o("el-dialog",{attrs:{title:"編輯學員",visible:t.editDialogVisible,width:"350px"},on:{"update:visible":function(e){t.editDialogVisible=e}}},[o("el-form",{ref:"form",attrs:{model:t.editForm,"label-width":"100px"}},[o("el-form-item",{attrs:{label:"狀態"}},[o("el-switch",{attrs:{"active-text":"報到","inactive-text":"未報到"},model:{value:t.editForm.status,callback:function(e){t.$set(t.editForm,"status",e)},expression:"editForm.status"}})],1),t._v(" "),o("el-form-item",{attrs:{label:"暱稱"}},[o("el-input",{model:{value:t.editForm.name,callback:function(e){t.$set(t.editForm,"name",e)},expression:"editForm.name"}})],1),t._v(" "),o("el-form-item",{attrs:{label:"組別"}},[o("el-input",{model:{value:t.editForm.group,callback:function(e){t.$set(t.editForm,"group",e)},expression:"editForm.group"}})],1),t._v(" "),o("el-form-item",{attrs:{label:"部門"}},[o("el-input",{model:{value:t.editForm.department,callback:function(e){t.$set(t.editForm,"department",e)},expression:"editForm.department"}})],1),t._v(" "),o("el-form-item",{attrs:{label:"員工編號"}},[o("el-input",{model:{value:t.editForm.employeeID,callback:function(e){t.$set(t.editForm,"employeeID",e)},expression:"editForm.employeeID"}})],1)],1),t._v(" "),o("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[o("el-button",{on:{click:function(e){t.editDialogVisible=!1}}},[t._v("取 消")]),t._v(" "),o("el-button",{attrs:{type:"primary"},on:{click:t.summitEditForm}},[t._v("送 出")])],1)],1)],1)],1)},r=[],s=(o("7f7f"),o("b6ce"),o("5441")),a=o.n(s),n=o("2ee3"),l={name:"Dashboard",data:function(){return{newForm:{status:!1,name:"",group:"",department:"",employeeID:""},editForm:{status:!1,name:"",group:"",department:"",employeeID:""},addDialogVisible:!1,editDialogVisible:!1,classID:0,list:[],listLoading:!0}},created:function(){var t=this;"undefined"!==typeof this.$route.query.classID?this.classID=this.$route.query.classID:a()({title:"課程ID錯誤",message:"錯誤訊息",confirmButtonText:"離開"}).then(function(e){t.$router.push({path:"/"})}),this.getStudentListFromApi()},methods:{goToScan:function(){this.$router.push({path:"/scan/index",query:{classID:this.classID,t:+new Date}})},getStudentListFromApi:function(){var t=this;Object(n["f"])(this.classID).then(function(e){e.error?(t.$message.error("取Api失敗"),a()("提示","Api失敗")):(t.list=e.data,t.listLoading=!1)})},enterStatus:function(t){this.editDialogVisible=!0,this.editForm.status=t.status,this.editForm.name=t.name,this.editForm.group=t.group,this.editForm.department=t.department,this.editForm.employeeID=t.employeeID},summitAddForm:function(){var t=this;this.newForm.employeeID=parseInt(this.newForm.employeeID,10),Object(n["i"])(this.classID,this.newForm).then(function(e){e.error?(t.$message.error("取Api失敗"),a()("提示","Api失敗")):(t.$message({message:"新增成功",type:"success"}),t.getStudentListFromApi(),t.addDialogVisible=!1)})},summitEditForm:function(){var t=this;this.editForm.employeeID=parseInt(this.editForm.employeeID,10),Object(n["d"])(this.classID,this.editForm).then(function(e){e.error?(t.$message.error("取Api失敗"),a()("提示","Api失敗")):(t.$message({message:"編輯成功",type:"success"}),t.getStudentListFromApi(),t.editDialogVisible=!1)})},deleteStudentByApi:function(t){var e=this;this.$confirm("此操作將永久刪除學員，是否繼續?","提示",{confirmButtonText:"確定",cancelButtonText:"取消",type:"warning"}).then(function(){Object(n["c"])(e.classID,t.employeeID).then(function(t){t.error?(e.$message.error("取Api失敗"),a()("提示","Api失敗")):(e.$message({message:"刪除成功",type:"success"}),e.getStudentListFromApi())})}).catch(function(){e.$message({type:"info",message:"已取消删除"})})}}},u=l,c=o("2877"),m=Object(c["a"])(u,i,r,!1,null,"374b2831",null);e["default"]=m.exports}}]);