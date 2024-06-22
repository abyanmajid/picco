"use client"

import React from "react";
import { Input } from "@nextui-org/react";
import { EyeFilledIcon } from "@/components/ui/icons";
import { EyeSlashFilledIcon } from "@/components/ui/icons";

export default function App({ includeConfirmPassword }: { includeConfirmPassword: boolean }) {
    const [isVisible, setIsVisible] = React.useState(false);
    const [confirmPasswordIsVisible, setConfirmPasswordIsVisible] = React.useState(false);

    const toggleVisibility = () => setIsVisible(!isVisible);
    const toggleConfirmPasswordVisibility = () => setConfirmPasswordIsVisible(!confirmPasswordIsVisible);

    return <>
        <Input
            label="Password"
            variant="bordered"
            placeholder="Enter your password"
            endContent={
                <button className="focus:outline-none" type="button" onClick={toggleVisibility}>
                    {isVisible ? (
                        <EyeSlashFilledIcon className="text-2xl text-default-400 pointer-events-none" />
                    ) : (
                        <EyeFilledIcon className="text-2xl text-default-400 pointer-events-none" />
                    )}
                </button>
            }
            type={isVisible ? "text" : "password"}
            className="max-w-md"
        />
        {includeConfirmPassword ?
            <Input
                label="Confirm Password"
                variant="bordered"
                placeholder="Re-enter your password"
                endContent={
                    <button className="focus:outline-none" type="button" onClick={toggleConfirmPasswordVisibility}>
                        {confirmPasswordIsVisible ? (
                            <EyeSlashFilledIcon className="text-2xl text-default-400 pointer-events-none" />
                        ) : (
                            <EyeFilledIcon className="text-2xl text-default-400 pointer-events-none" />
                        )}
                    </button>
                }
                type={confirmPasswordIsVisible ? "text" : "password"}
                className="max-w-md"
            /> : ""
        }
    </>;
}