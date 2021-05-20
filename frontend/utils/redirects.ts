interface RedirectReturn {
  redirect: {
    permanent: boolean
    destination: string
  }
}

export const redirectToDashboard = (): RedirectReturn => {
  return {
    redirect: {
      permanent: false,
      destination: '/dashboard',
    },
  }
}

export const redirectToLandingPage = (): RedirectReturn => {
  return {
    redirect: {
      permanent: false,
      destination: '/',
    },
  }
}

export const redirectToErrorPage = (message?: string): RedirectReturn => {
  return {
    redirect: {
      permanent: false,
      destination: `/error${message ? `?message=${message}` : ''}`,
    },
  }
}
