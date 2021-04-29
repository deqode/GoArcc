import React, { useContext } from 'react'
import { Router, useRouter } from 'next/router';
import { SERVER } from '../../utils/constants';
import { UserContext } from '../../Contexts/UserContext';

function callback() {
    const { setUser } = useContext(UserContext)
    const router = useRouter();
    const { query } = router
    const authenticate = async () => {
        let res = await fetch(`${SERVER}/authentication/callback?code=${query.code}&state=${query.state}`)
        let data = await res.json()
        if (data.idToken)
            setUser(data)
        router.push("/")
    }
    authenticate()
    return (
        <div>

        </div>
    )
}

export default callback

