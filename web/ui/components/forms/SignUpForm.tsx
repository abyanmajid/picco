import Container from "@/components/common/Container"
import EmailInput from "@/components/auth/EmailInput"
import PasswordInput from "@/components/auth/PasswordInput"
import { CredentialSignUpButton, OAuthSignUpButton, AuthRedirectButton } from "@/components/auth/AuthButtons"
import { Divider } from "@nextui-org/divider"

export default function SignUpForm() {
    return <>
        <Container className="w-full max-w-md grid gap-4 mb-2">
            <EmailInput />
            <PasswordInput includeConfirmPassword={true} />
            <CredentialSignUpButton />
            <Divider className="my-2" />
            <OAuthSignUpButton provider="google" />
            <OAuthSignUpButton provider="github" />
        </Container>
        <AuthRedirectButton redirectTo="/login" />
    </>
}