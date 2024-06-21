import Container from "@/components/common/Container"
import EmailInput from "@/components/auth/EmailInput"
import { AuthRedirectButton, SendPasswordResetEmailButton } from "@/components/auth/AuthButtons"
import { Link } from "@nextui-org/link"

export default function ForgotPasswordForm() {
    return <>
        <Container className="w-full max-w-md grid gap-4 mb-2">
            <EmailInput />
            <SendPasswordResetEmailButton />
        </Container>
        <Link href="/login" className="text-neutral-400">Go back</Link>
    </>
}