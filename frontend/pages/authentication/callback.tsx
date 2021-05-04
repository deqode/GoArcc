import React, { useContext, useEffect } from 'react'
import { Router, useRouter } from 'next/router';
import { SERVER, CENTRIFUGO } from '../../utils/constants';
import Centrifuge from 'centrifuge'
import { UserContext } from '../../Contexts/UserContext';
import { AppContext } from '../../Contexts/AppContext';
function callback() {
    const { setUser } = useContext(UserContext)
    const { app, setApp } = useContext(AppContext)
    const router = useRouter();
    const { query } = router
    useEffect(() => {
        (async () => {
            if (query.code) {
                console.log("send", query)
                let res = await fetch(`${SERVER}/authentication/callback?code=${query.code}&state=${query.state}`)
                let data = await res.json()
                if (data.idToken) {
                    let user = data
                    user.state = 1
                    setUser(user)
                    const centrifuge = new Centrifuge(CENTRIFUGO)
                    centrifuge.connect();
                    setApp({centrifuge, subscribe: false})
                }
                router.push("/")
            }
        })()
    }, [router, query])

    // useEffect(() => {
    //     app.centrifuge.subscribe('clone-logs', (message) => {
    //         console.log(message)
    //     })
    // }, [app])
    return (
        <div>
        </div>
    )
}
export default callback