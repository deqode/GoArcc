import React from 'react'
import { Router, useRouter } from 'next/router';
import { SERVER } from '../../utils/constants';
import { token } from '../../GraphQL/cache';

function callback() {
    const router = useRouter();
    const { query } = router
    const authenticate = async () => {
        let res = await fetch(`${SERVER}/authentication/callback?code=${query.code}&state=${query.state}`)
        let data = await res.json()
        if (data.idToken)
            token(data)
        router.push("/")
    }
    authenticate()
    return (
        <div>

        </div>
    )
}

export default callback

