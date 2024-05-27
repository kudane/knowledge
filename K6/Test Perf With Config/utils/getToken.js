import http from 'k6/http';

export default function getToken() {
    const body = {};

    const response = http.post('http://localhost:3000/login', body);

    return response.json().access_token;
}