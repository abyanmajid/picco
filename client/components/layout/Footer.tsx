import { Link } from "@nextui-org/link";

export default function Footer() {
    return <footer className="w-full flex items-center justify-center pt-3">
        <Link
            isExternal
            className="flex items-center gap-1 text-current"
            href="https://github.com/abyanmajid"
            title="https://github.com/abyanmajid"
        >
            <span className="text-default-600">Built by</span>
            <p className="text-primary">Abyan Majid</p>
        </Link>
    </footer>
}