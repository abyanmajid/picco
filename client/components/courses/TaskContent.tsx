import { Button } from "@nextui-org/button"
import Container from "../common/Container"
import Title from "../common/Title"
import TaskContentHeader from "./TaskContentHeader"
import TaskInstructions from "@/content/introduction-to-programming/module-1/task-1.mdx"
import { MDXRemote } from 'next-mdx-remote/rsc'

export default function TaskContent() {
    const markdown = `
# Multi-line MDX Content

This is a paragraph with some **bold** text and some _italic_ text.

Here is an example of inline code: \`const x = 10;\`

## Code Block Example

Here is a code block:

\`\`\`js
function greet(name) {
  return \`Hello, \${name}!\`;
}

console.log(greet('World'));
\`\`\`

Another paragraph with more information. You can add as many lines as needed.
`
    return <>
        <TaskContentHeader />
        <Container className="prose dark:prose-invert">
            <MDXRemote source={markdown} />
        </Container>
    </>
}