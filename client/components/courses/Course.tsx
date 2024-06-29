"use client";

import React from "react";
import { Card, CardBody } from "@nextui-org/react";
import { Link } from "@nextui-org/react";
import Container from "../common/Container";
import { Divider } from "@nextui-org/react"
import { Button } from "@nextui-org/react";
import { Avatar } from "@nextui-org/react";
import { Accordion, AccordionItem } from "@nextui-org/react";
import { IoMdCheckmarkCircleOutline as CheckmarkIcon } from "react-icons/io";
import { Chip } from "@nextui-org/react";
import SideButtons from "./SideButtons"

import Title from "../common/Title";
import Subtitle from "../common/Subtitle";
import CourseMetrics from "./CourseMetrics";

type courseDetails = {
    title: string,
    description: string,
    creator: string,
    creatorProfileUrl: string,
    modulesURL: string,
    likes: number,
}

type Props = {
    courseDetails: courseDetails,
}

export default function Course({ courseDetails }: Props) {
    
    return <>
        <Card className="bg-black border border-neutral-700 w-full text-left p-8">
            <CardBody>
                <Container className="grid grid-cols-2">
                    <Title size="sm">{courseDetails.title}</Title>
                    <SideButtons/> 
                </Container>
                <Subtitle className="my-4">{courseDetails.description}</Subtitle>
                <Container className="grid grid-cols-2">
                   <CourseMetrics creator={courseDetails.creator} creatorProfileUrl={courseDetails.creatorProfileUrl}/> 
                </Container>
                <Divider className="my-4" />
                <Accordion>
                    <AccordionItem key="1" subtitle="Press to expand" title="Topics (3)">
                        <ol className="list-decimal list-inside mb-4">
                            <li><span className="font-semibold">Variables and datatypes:</span> Lorem ipsum dolor Lorem ipsum dolor</li>
                            <li><span className="font-semibold">Functions:</span> Lorem ipsum dolor Lorem ipsum dolor</li>
                            <li><span className="font-semibold">Classes and objects:</span> Lorem ipsum dolor</li>
                        </ol>
                    </AccordionItem>
                    <AccordionItem key="2" subtitle="Press to expand" title="Modules (26)">
                        <Accordion variant="splitted">
                            <AccordionItem key="1" title={<span className="text-success flex items-center"><CheckmarkIcon size={20} className="mr-2" />Variables and datatypes</span>} className="text-neutral-300">
                                <ol className="list-decimal list-inside mb-4">
                                    <li className="text-success"><span className="font-semibold">Lecture:</span> Defining abstract data types in Python <Chip color="success" variant="flat">Completed</Chip></li>
                                    <li><span className="font-semibold">Quiz:</span> Debugging variables and data type errors</li>
                                    <li><span className="font-semibold">Exercise:</span> Leap year</li>
                                    <li><span className="font-semibold">Guided project:</span> Building a web server with Flask</li>
                                </ol>
                                <Button variant="shadow" className="mb-2">Go to module</Button>
                            </AccordionItem>
                            <AccordionItem key="2" title="Classes and objects">
                                <ol className="list-decimal list-inside mb-4">
                                    <li><span className="font-semibold">Variables and datatypes:</span> Lorem ipsum dolor Lorem ipsum dolor</li>
                                </ol>
                                <Button variant="shadow" className="mb-2">Go to module</Button>
                            </AccordionItem>
                            <AccordionItem key="3" title="GUIDED PROJECT: Building a web server in Flask">
                                <ol className="list-decimal list-inside mb-4">
                                    <li><span className="font-semibold">Variables and datatypes:</span> Lorem ipsum dolor Lorem ipsum dolor</li>
                                </ol>
                                <Button variant="shadow" className="mb-2">Go to module</Button>
                            </AccordionItem>
                            <AccordionItem key="3" title="GUIDED PROJECT: Building a web server in Flask">
                                <ol className="list-decimal list-inside mb-4">
                                    <li><span className="font-semibold">Variables and datatypes:</span> Lorem ipsum dolor Lorem ipsum dolor</li>
                                </ol>
                                <Button variant="shadow" className="mb-2">Go to module</Button>
                            </AccordionItem>
                            <AccordionItem key="3" title="GUIDED PROJECT: Building a web server in Flask">
                                <ol className="list-decimal list-inside mb-4">
                                    <li><span className="font-semibold">Variables and datatypes:</span> Lorem ipsum dolor Lorem ipsum dolor</li>
                                </ol>
                                <Button variant="shadow" className="mb-2">Go to module</Button>
                            </AccordionItem>
                            <AccordionItem key="3" title="GUIDED PROJECT: Building a web server in Flask">
                                <ol className="list-decimal list-inside mb-4">
                                    <li><span className="font-semibold">Variables and datatypes:</span> Lorem ipsum dolor Lorem ipsum dolor</li>
                                </ol>
                                <Button variant="shadow" className="mb-2">Go to module</Button>
                            </AccordionItem>
                        </Accordion>
                    </AccordionItem>
                </Accordion>
            </CardBody>
        </Card >
    </>
}