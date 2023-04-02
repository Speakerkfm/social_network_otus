import http from 'k6/http'
import {check, sleep} from 'k6'

export let options = {
    discardResponseBodies: false,
    scenarios: {
        contacts: {
            executor: "ramping-vus",
            startVUs: 1,
            stages: [
                {duration: "1m", target: 10},
                {duration: "1m", target: 100},
                {duration: "1m", target: 1000},
            ],
            gracefulRampDown: "0s",
        },
    },
};

export default function () {
    let res = http.post('http://127.0.0.1/user/register', `{
        "first_name": "Имя",
        "second_name": "Фамилия",
        "age": 18,
        "biography": "Хобби, интересы и т.п.",
        "city": "Москва",
        "sex": "male",
        "password": "Секретная строка"
    }`, {
        headers: { 'Content-Type': 'application/json' },
    });

    check(res, {'User register': (r) => r.status === 200})

    res = http.get('http://127.0.0.1/user/get/'+res.json().user_id)

    check(res, {'User get': (r) => r.status === 200})

    sleep(1)
}
