import React, { useContext, useEffect } from 'react'
import { useRouter } from 'next/router';
import { SERVER, CENTRIFUGO } from '../../utils/constants';
import Centrifuge from 'centrifuge'
import { UserContext } from '../../Contexts/UserContext';
import { AppContext } from '../../Contexts/AppContext';
export default function Callback() {
    const { setUser } = useContext(UserContext)
    const { setApp } = useContext(AppContext)
    const router = useRouter();
    const { query } = router
    useEffect(() => {
        (async () => {
            if (query.code) {
                console.log("send", query)
                const res = await fetch(`${SERVER}/authentication/callback?code=${query.code}&state=${query.state}`)
                const data = await res.json()
                if (data.idToken) {
                    const user = data
                    user.state = 1
                    setUser(user)
                    const centrifuge = new Centrifuge(CENTRIFUGO)
                    centrifuge.connect();
                    setApp({ centrifuge, subscribe: false })
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
