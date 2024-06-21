import React from "react"
import { Input } from "@nextui-org/react"

export default function EmailInput() {
    return <Input type="email" variant="bordered" label="Email Address" placeholder="Enter your email address..." className="w-full" />
}