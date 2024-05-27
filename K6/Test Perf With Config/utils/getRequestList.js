export default function getRequestList(accessToken, body) {
    /* your can custom `body` with 
    body.somthing = 777;
    ... */

    return [
        'POST',
        'http://localhost:3000/create/something',
        JSON.stringify(body),
        {
            headers: {
                'Authorization': `Bearer ${accessToken}`,
                'Content-Type': 'application/json',
            }
        }
    ];
}