import * as config from "../auth_config.json"

const { domain, clientId, authorizationParams: { audience }, apiUri, appUri, errorPath } = config as {
    domain: string,
    clientId: string,
    authorizationParams: {
        audience?: string
    },
    apiUri: string,
    appUri: string,
    errorPath: string
}

export const environment = {
    production: false,
    auth: {
        domain,
        clientId,
        authorizationParams: {
            ...(audience && audience !== 'YOUR_API_IDENTIFIER' ? { audience } : null),
            redirect_uri: "http://localhost:4200",
        },
        errorPath,
    },
    httpInterceptor: {
        allowedList: [`${apiUri}/*`]
    }
}