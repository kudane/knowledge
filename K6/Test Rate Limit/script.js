import http from "k6/http";
import { Counter } from "k6/metrics";

export let options = {
    vus: 1,
    duration: '5s',
    iterations: 10,
};

let CounterSuccess = new Counter("Success");

let CounterOverLimit = new Counter("OverLimit");

export default function () {
    const res = http.get('https://localhost:3000/test/rate-limit');

    switch (res.status) {
        case 200:
            CounterSuccess.add(1);
            break;
        default:
            CounterOverLimit.add(1);
            break;
    }
};