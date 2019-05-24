import axios from 'axios';

/**
 * Create Axios
 */
export const http = axios.create({
    baseURL: "localhost:9090",
    headers: {'Content-Type': 'application/x-www-form-urlencoded'}
});


/**
 * Handle all error messages.
 */
http.interceptors.response.use(function (response) {
    return response;
}, function (error) {
    const { response } = error;

    if ([401].indexOf(response.status) >= 0) {
        if (response.status == 401 && response.data.error.message != 'Unauthorized') {
            return Promise.reject(response);
        }
        window.location = '/login';
    }

    if([422].indexOf(response.status) >= 0){
        let message = response.message;
        for(var i in response.data.errors){
            message = response.data.errors[i][0];
        }
        alert(message);
    }

    return Promise.reject(error);
});

export default function install(Vue) {
    Object.defineProperty(Vue.prototype, '$http', {
        get() {
            return http
        }
    })
}