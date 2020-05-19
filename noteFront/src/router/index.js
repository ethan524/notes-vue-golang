import Vue from 'vue'
import Router from 'vue-router'
import Login from '../components/Login';
import Register from '../components/Register';
import Home from '../components/Home';
import Index from '../components/Index';
import Follow from '../components/Follow';
import Comment from '../components/Comment';
import Likes from '../components/Likes';

Vue.use(Router)

export default new Router({
  routes: [
    {
      path:'/login',
      name:'Login',
      meta: { title: '登录' },
      component:Login
    },
    {
      path:'/register',
      name:'Register',
      meta: { title: '注册' },
      component:Register
    },
    {
      path:'/',
      name:'Home',
      component:Home,
      redirect:'/index',
      meta: { title: '首页' },
      children:[
        {
          path:'/index',
          name:"Index",
          meta: { title: '首页' },
          component:Index,
        },
        {
          path:'/follow',
          meta: { title: '关注' },
          component:Follow
        },
        {
          path:'/comment',
          meta: { title: '评论' },
          component:Comment
        },
        {
          path:'/likes',
          meta: { title: '喜欢和赞' },
          component:Likes
        }
      ]
    },
    
  ]
})
