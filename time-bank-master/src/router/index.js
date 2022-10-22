import VueRouter from "vue-router";
import main from "@/components/main";
import Login from "@/components/login/Login";
import Register from "@/components/Register/Register"
import SearchServed from "@/components/search/SearchServed"
import SearchServicing from "@/components/search/SearchServicing"
import VipInfo from "@/components/info/VipInfo";
import SearchRechargeList from "@/components/search/SearchRechargeList";
import SearchSpecialTxList from "@/components/search/SearchSpecialTxList";
//创建路由
const router= new VueRouter({
    routes:[
        { path: '/', redirect: '/main' },
        {
            path:'/main',
            name:'main',
            component:main
        },
        { path:'/Login',name: 'Login',component:Login},
        { path:'/Register',name: 'Register',component:Register},
        { path:'/SearchServed',name: 'SearchServed',component:SearchServed},
        { path:'/SearchServicing',name: 'SearchServicing',component:SearchServicing},
        { path:'/VipInfo',name: 'VipInfo' , component: VipInfo},
        { path:'/SearchRechargeList',name: 'SearchRechargeList',component:SearchRechargeList},
        { path:'/SearchSpecialTxList',name: 'SearchSpecialTxList',component:SearchSpecialTxList}
    ]
});
//导出路由

export default router
