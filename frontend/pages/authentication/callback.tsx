import React, { useContext, useEffect } from 'react'
import { useRouter } from 'next/router';
import { SERVER, CENTRIFUGO } from '../../utils/constants';
import Centrifuge from 'centrifuge'
import { AppContext } from '../../Contexts/AppContext';
export default function Callback() {
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
                    const sessionRes = await fetch(`/api/session/set`, {
                        method: 'POST',
                        headers: {
                            'Accept': 'application/json',
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify(data)
                    })
                    const sessionData = await sessionRes.json()
                    if (!sessionData) {
                        router.push("/dashboard")
                        return
                    }

                    const centrifuge = new Centrifuge(CENTRIFUGO)
                    centrifuge.connect();
                    setApp({ centrifuge, subscribe: false })
                }
                router.push("/dashboard")
            }
        })()
    }, [router, query, setApp])

    return (
        <div>
        </div>
    )
}
