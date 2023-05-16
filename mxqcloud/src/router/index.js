import {createRouter, createWebHashHistory} from 'vue-router'
import Upload from '@/views/Upload.vue'
import Download from '@/views/Download.vue'
import Chat from '@/views/Chat.vue'
import Mine from '@/views/Mine.vue'

const routes = [
    {
        path: '/',
        name: 'upload',
        component: Upload
    },
    {
        path: '/download',
        name: 'download',
        component: Download
    },
    {
        path: '/chat',
        name: 'chat',
        component: Chat
    },
    {
        path: '/mine',
        name: 'mine',
        component: Mine
    },
]
const router = createRouter({
    history: createWebHashHistory(process.env.BASE_URL),
    routes
})
export default router
