import React from "react";
import { Card, CardBody } from "@nextui-org/react";
import { Link } from "@nextui-org/react";
import Container from "../common/Container";
import { Divider } from "@nextui-org/react"
import { Button } from "@nextui-org/react";
import { HeartIcon } from "../ui/icons";
import { Avatar } from "@nextui-org/react";
import { Accordion, AccordionItem } from "@nextui-org/react";
import { capitalize } from "@/utils/helpers";

import Title from "../common/Title";
import Subtitle from "../common/Subtitle";
import { IoMdCheckmarkCircleOutline as CheckmarkIcon } from "react-icons/io";


type Props = {
    taskName: string,
    taskType: "lecture" | "exercise" | "quiz",
    isCompleted: boolean,
}

export default function TaskCard({ taskName, taskType, isCompleted }: Props) {
    return <>
        <Link href="#">
            <Card className={`bg-black border w-full text-left p-4 hover:bg-neutral-900 mb-4 ${isCompleted ? "text-success border-green-400" : ""}`}>
                <CardBody>
                    <Container className="grid grid-cols-2">
                        <Container className="flex items-center gap-6">
                            <CheckmarkIcon size={48} />
                            <div className="flex flex-col">
                                <div className="text-xl font-semibold">{taskName}</div>
                                <div className="text-md text-neutral-400">{capitalize(taskType)}</div>
                            </div>
                        </Container>
                        <Container className="justify-end text-right font-bold">
                            {isCompleted ? "Completed" : ""}
                        </Container>
                    </Container>
                </CardBody>
            </Card >
        </Link>
    </>
}