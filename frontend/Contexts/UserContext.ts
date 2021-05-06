import { createContext } from "react";
import { UserContextInterface } from "../interface";

const userContext: UserContextInterface = {
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
    setUser: () => null,
    removeUser: () => null


}

export const UserContext = createContext<UserContextInterface>(userContext);
