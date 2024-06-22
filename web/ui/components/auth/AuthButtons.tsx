import { Button } from "@nextui-org/button"
import { capitalize } from "@/utils/helpers"
import { FcGoogle as GoogleIcon } from "react-icons/fc";
import { GithubIcon } from "../ui/icons";
import { Link } from "@nextui-org/link";

type OAuthButtonProps = {
    provider: "google" | "github"
}

export function CredentialLoginButton() {
    return <>
        <Button color="primary" variant="solid" className="w-full">
            Log In
        </Button>
    </>
}

export function CredentialSignUpButton() {
    return <>
        <Button color="primary" variant="solid" className="w-full">
            Sign Up
        </Button>
    </>
}

export function OAuthLoginButton({ provider }: OAuthButtonProps) {
    const oAuthProvider = capitalize(provider)
    return <>
        <Button color="default" variant="ghost" className="w-full">
            {provider === "google" ?
                <>
                    <GoogleIcon size={24} /> &nbsp;Continue in with {oAuthProvider}
                </>
                :
                <>
                    <GithubIcon /> &nbsp;Continue in with {oAuthProvider}
                </>
            }
        </Button>
    </>
}
export function OAuthSignUpButton({ provider }: OAuthButtonProps) {
    const oAuthProvider = capitalize(provider)
    return <>
        <Button color="default" variant="ghost" className="w-full">
            {provider === "google" ?
                <>
                    <GoogleIcon size={24} /> &nbsp;Sign up with {oAuthProvider}
                </>
                :
                <>
                    <GithubIcon /> &nbsp;Sign up with {oAuthProvider}
                </>
            }
        </Button>
    </>
}

export function AuthRedirectButton({ redirectTo }: { redirectTo: "/signup" | "/login" }) {
    return <p>
        {redirectTo === "/signup" ? "New to codemore.io?" : "Already have an account?"}&nbsp;
        <Link href={redirectTo} className="text-primary">
            {redirectTo === "/signup" ? "Sign Up" : "Login"}
        </Link>
    </p>
}

export function ForgotPassword() {
    return <p>
        <Link href="/forgot-password" className="text-neutral-400">
            Forgot Password?
        </Link>
    </p>
}

export function SendPasswordResetEmailButton() {
    return <>
        <Button color="primary" variant="solid" className="w-full">
            Send Reset Link
        </Button>
    </>
}

export function ResetPasswordButton() {
    return <>
        <Button color="primary" variant="solid" className="w-full">
            Reset Password
        </Button>
    </>
}