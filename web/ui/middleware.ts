import { withMiddlewareAuthRequired } from '@auth0/nextjs-auth0/edge';

export default withMiddlewareAuthRequired();

export const config = {
    matcher: [
        '/courses/:courseId/modules/:path*',
        '/courses/:courseId/modules/:moduleId/tasks/:path*',
        '/admin',
        '/bug-hunt'
    ]
};