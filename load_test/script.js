import http from 'k6/http'
import {check, sleep} from 'k6'

export let options = {
    discardResponseBodies: true,
    scenarios: {
        contacts: {
            executor: "ramping-vus",
            startVUs: 1,
            stages: [
                { duration: "1m", target: 10 },
                { duration: "1m", target: 100 },
                { duration: "1m", target: 1000 },
            ],
            gracefulRampDown: "0s",
        },
    },
};

export default function () {
    let res = http.get('http://127.0.0.1/user/search?first_name=Константин&last_name=А')

    check(res, {'User search': (r) => r.status === 200})

    sleep(1)
}
