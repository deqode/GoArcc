import { createContext } from "react";

export const UserContext = createContext({
    user: {
        accessToken: "",
        idToken: "",
        userId: "",
        provider: "",
        accounts: [],
        state: 0,
        //  loading=0
        //  authenticated=1
        //  unauthenticated=-1
    },
    setUser: (value) => { },
    removeUser: () => { }
});
