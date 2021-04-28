import React, { useContext, useEffect } from 'react'
import { Router, useRouter } from 'next/router';

function callback() {
    const router = useRouter();
    const { query } = router
    useEffect(() => {
        console.log(query)
        if (query.code && query.code.length > 0){
            console.log(query)
            // router.push("/tell-us-more")
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

