import Container from "@/components/common/Container"
import EmailInput from "@/components/auth/EmailInput"
import PasswordInput from "@/components/auth/PasswordInput"
import { CredentialLoginButton, OAuthLoginButton, AuthRedirectButton, ForgotPassword } from "@/components/auth/AuthButtons"
import { Divider } from "@nextui-org/divider"

export default function LoginForm() {
    return <>
        <Container className="w-full max-w-md grid gap-4 mb-2">
            <EmailInput />
            <PasswordInput includeConfirmPassword={false} />
            <CredentialLoginButton />
            <ForgotPassword />
            <Divider className="my-2" />
            <OAuthLoginButton provider="google" />
            <OAuthLoginButton provider="github" />
        </Container>
        <AuthRedirectButton redirectTo="/signup" />
    </>
}