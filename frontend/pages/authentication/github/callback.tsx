import React, { useContext, useEffect } from 'react'
import { Router, useRouter } from 'next/router';
import { UserContext } from '../../../Contexts/UserContext';
import { SERVER } from '../../../utils/constants';

function callback() {
    const router = useRouter();
    const { setUser, user } = useContext(UserContext)

    const { query } = router
    useEffect(() => {
        if (query.code && query.code.length > 0) {
            (async () => {
                console.log(query)
                if (query.code) {
                    let res = await fetch(`${SERVER}/vcs-connection/GITHUB/callback?code=${query.code}&account_id=7cd96426-5f94-1a78-e9b2-b96bf41e1a45`, {
                        headers: new Headers({
                            'Authorization': `Bearer ${user.idToken}`,
                        })
                    })
                    let data = await res.json()
                    let newUser = user
                    newUser.provider = data.provider
                    console.log(newUser)
                    setUser(newUser)
                    router.push("/tell-us-more")
                } else {
                    router.push("/")
                }

            })()

        }
        // else
        // router.push("/")
    }, [router])

    return (
        <div>

        </div>
    )
}

export default callback

