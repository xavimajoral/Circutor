import * as jwt from 'react-jwt';

export async function LoginUser(_email, _pass) {

    const response = await fetch(`http://localhost:1234/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email: _email, password: _pass })
    })

    return await response.json();

}

export async function RegisterUser(_email, _pass) {

    const response = await fetch(`http://localhost:1234/signup`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email: _email, password: _pass })
    })

    return await response.json();

}

export function TokenIsValid(_token) {
    const myDecodedToken = jwt.decodeToken(_token);

    if (jwt.isExpired(_token)) {
        return false
    }
    return true
}