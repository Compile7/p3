import { useCallback, useState } from 'react'
import './auth.css'
import { User } from '../User' // component display user (see detail on /example directory)
import {LoginSocialGoogle} from 'reactjs-social-login'

// CUSTOMIZE ANY UI BUTTON
import {
//   FacebookLoginButton,
  GoogleLoginButton,
//   GithubLoginButton,
//   AmazonLoginButton,
//   InstagramLoginButton,
//   LinkedInLoginButton,
//   MicrosoftLoginButton,
//   TwitterLoginButton,
//   AppleLoginButton,
} from 'react-social-login-buttons'
import { IdTokenClient } from 'google-auth-library'

// import { ReactComponent as PinterestLogo } from './assets/pinterest.svg'
// import { ReactComponent as TiktokLogo } from './assets/tiktok.svg'

// REDIRECT URL must be same with URL where the (reactjs-social-login) components is locate
// MAKE SURE the (reactjs-social-login) components aren't unmounted or destroyed before the ask permission dialog closes
// const REDIRECT_URI = window.location.href;

const Auth = () => {
  const [provider, setProvider] = useState('')
  const [profile, setProfile] = useState<any>()

  const onLoginStart = useCallback(() => {
    alert('login start')
  }, [])

  const onLogoutSuccess = useCallback(() => {
    setProfile(null)
    setProvider('')
    alert('logout success')
  }, [])

  return (
    <>
      {provider && profile ? (
        <User provider={provider} profile={profile} onLogout={onLogoutSuccess} />
      ) : (
        <div className={`App ${provider && profile ? 'hide' : ''}`}>
          <h1 className='title'>ReactJS Social Login</h1>

          
          <LoginSocialGoogle
            isOnlyGetToken
            typeResponse="idToken"
            access_type='offline'
            prompt='consent'
            client_id={import.meta.env.VITE_APP_GG_APP_ID || ''}
            onLoginStart={onLoginStart}
            onResolve={({ provider, data }: any) => {
              console.log(provider)
              console.log(data)
              setProvider(provider)
              setProfile(data)
            }}
            onReject={(err: any) => {
              console.log(err)
            }}
          >
            <GoogleLoginButton />
          </LoginSocialGoogle>

        </div>
      )}
    </>
  )
}

export default Auth