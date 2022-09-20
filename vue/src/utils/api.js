import axios from '../utils/http'
import QS from 'qs'

export function loginRequest(name,password) {
    return axios({
        url: `/api/login`,
        method: 'post',
        data: QS.stringify({name:name,password:password})
    })
}

export function getProfileRequest() {
    return axios({
        url: `/api/profile`,
        method: 'get'
    })
}

export function putProfileNicknameRequest(nickname) {
    return axios({
        url: `/api/profile/nickname`,
        method: 'put',
        data: QS.stringify({nickname:nickname})
    })
}

export function putProfilePictureRequest(form) {
    return axios({
        url: `/api/profile/picture`,
        method: 'put',
        form:form
    })
}