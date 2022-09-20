import axios from 'axios'
import { Loading, Message } from 'element-ui'
import router from '../router/index.js'

let loading

function startLoading() {
    loading = Loading.service({
        lock: true,
        text: '加载中....',
        background: 'rgba(0, 0, 0, 0.7)'
    })
}

function endLoading() {
    loading.close()
}

// 请求拦截
axios.interceptors.request.use(
    (confing) => {
        startLoading()
        //设置请求头
        if (localStorage.eToken) {   
            confing.headers.Authorization = localStorage.eToken
        }
        return confing
    },
    (error) => {
        endLoading()
        return Promise.reject(error)
    }
)



//响应拦截

axios.interceptors.response.use(
    (response) => {
        endLoading()
        if (response.data.code !== 1) {
            if (response.data.code === -301) {
                if (localStorage.getItem('eToken')){
                    Message.error('请重新登录')
                    localStorage.removeItem('eToken')
                }
                router.replace({name:'login'}).then(() => {
                    return response;
                })
            } else {
                Message.error(response.data.msg)
            }
        }
        return response
    },
    (error) => {
        endLoading()
        Message.error("请求错误")
        return Promise.reject(error)
    }
)
export default axios