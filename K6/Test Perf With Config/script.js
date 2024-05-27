import http from 'k6/http';
import { check } from 'k6';
import { getToken } from '../utils/getToken.js';
import { getRequestList } from '../utils/getRequestList.js';

// ใช้คำสั่ง `open('...')` ได้เฉพาะไฟล์ที่รัน k6 เท่านั้น, ให้ส่งไปเป็น parameter แทน
const body = JSON.parse(open('../json/body.json'));

export default function () {
    const accessToken = getToken(config.username);

    const listOfRequest = getRequestList(accessToken, body);

    const responses = http.batch(listOfRequest);

    if (responses && responses.length) {
        for (var i = 0; i < responses.length; i++) {
            check(responses[i], {
                'status was 200': (res) => res.status === 200,
            });
        }
    }
}