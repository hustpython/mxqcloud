<template>
    <Header/>
    <div class="upload">

        <n-upload
                style="padding-top: 10px;"
                multiple
                list-type="image-card"
                v-model:file-list="fileList"
                :action=apiServer
                :headers="{
      'device': res
    }"
                :custom-request="customRequest"
        >
            点击上传

        </n-upload>
        <n-divider/>
        <div class="uploadedList">
            <n-collapse arrow-placement="right">
                <n-collapse-item title="已完成" name="1">
                    <n-card title="2023-5-16" size="small"
                    >
                        <template #header-extra>
                            <n-checkbox size="small">
                                全选
                            </n-checkbox>
                        </template>
                        <div completeDatail>
                            <n-space justify="space-between">
                                <n-image
                                        width=51
                                        src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"/>
                                <n-space vertical>
                                    <p>
                                        vue.svg
                                    </p>
                                    <p>
                                        2023-05-16 12:00 301.0M
                                    </p>
                                </n-space>
                                <n-checkbox size="small"/>
                            </n-space>
                        </div>
                        <div completeDatail>
                            <n-space justify="space-between">
                                <n-image
                                        width=51
                                        src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"/>
                                <n-space vertical>
                                    <p>
                                        vue.svg
                                    </p>
                                    <p>
                                        2023-05-16 12:00 301.0M
                                    </p>
                                </n-space>
                                <n-checkbox size="small"/>
                            </n-space>
                        </div>
                        <div completeDatail>
                            <n-space justify="space-between">
                                <n-image
                                        width=51
                                        src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"/>
                                <n-space vertical>
                                    <p>
                                        vue.svg
                                    </p>
                                    <p>
                                        2023-05-16 12:00 301.0M
                                    </p>
                                </n-space>
                                <n-checkbox size="small"/>
                            </n-space>
                        </div>
                        <div completeDatail>
                            <n-space justify="space-between">
                                <n-image
                                        width=51
                                        src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"/>
                                <n-space vertical>
                                    <p>
                                        vue.svg
                                    </p>
                                    <p>
                                        2023-05-16 12:00 301.0M
                                    </p>
                                </n-space>
                                <n-checkbox size="small"/>
                            </n-space>
                        </div>
                        <div completeDatail>
                            <n-space justify="space-between">
                                <n-image
                                        width=51
                                        src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"/>
                                <n-space vertical>
                                    <p>
                                        vue.svg
                                    </p>
                                    <p>
                                        2023-05-16 12:00 301.0M
                                    </p>
                                </n-space>
                                <n-checkbox size="small"/>
                            </n-space>
                        </div>
                        <div completeDatail>
                            <n-space justify="space-between">
                                <n-image
                                        width=51
                                        src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"/>
                                <n-space vertical>
                                    <p>
                                        vue.svg
                                    </p>
                                    <p>
                                        2023-05-16 12:00 301.0M
                                    </p>
                                </n-space>
                                <n-checkbox size="small"/>
                            </n-space>
                        </div>
                    </n-card>

                    <n-card title="2023-5-16" size="small">
                        <div completeDatail>
                            <n-space justify="space-between">
                                <n-image
                                        width=51
                                        src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"/>
                                <n-space vertical>
                                    <p>
                                        vue.svg
                                    </p>
                                    <p>
                                        2023-05-16 12:00 301.0M
                                    </p>
                                </n-space>
                                <n-checkbox size="small"/>
                            </n-space>
                        </div>
                        <div completeDatail>
                            <n-space justify="space-between">
                                <n-image
                                        width=51
                                        src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"/>
                                <n-space vertical>
                                    <p>
                                        vue.svg
                                    </p>
                                    <p>
                                        2023-05-16 12:00 301.0M
                                    </p>
                                </n-space>
                                <n-checkbox size="small"/>
                            </n-space>
                        </div>
                        <div completeDatail>
                            <n-space justify="space-between">
                                <n-image
                                        width=51
                                        src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"/>
                                <n-space vertical>
                                    <p>
                                        vue.svg
                                    </p>
                                    <p>
                                        2023-05-16 12:00 301.0M
                                    </p>
                                </n-space>
                                <n-checkbox size="small"/>
                            </n-space>
                        </div>
                    </n-card>

                </n-collapse-item>
            </n-collapse>

        </div>
    </div>
    <Footer/>
</template>

<script setup>
    import Header from '../components/Header.vue'
    import Footer from '../components/Footer.vue'
    import {CloudUploadSharp} from "@vicons/material";
    import SERVER_API from '@/util/config.js'
    import {lyla} from "lyla";
    import {ref} from 'vue'
    import {GetDevice} from '@/util/Device.js'

    let res = GetDevice();

    const customRequest = ({
                               file,
                               data,
                               headers,
                               withCredentials,
                               action,
                               onFinish,
                               onError,
                               onProgress
                           }) => {
        const formData = new FormData();
        if (data) {
            Object.keys(data).forEach((key) => {
                formData.append(
                    key,
                    data[key]
                );
            });
        }
        formData.append('file', file.file);
        lyla.post(action, {
            withCredentials,
            headers,
            body: formData,
            onUploadProgress: ({percent}) => {
                onProgress({percent: Math.ceil(percent)});
            }
        }).then(({json}) => {
            console.log(JSON.stringify(json));
            onFinish()
        }).catch((error) => {
            console.log(error.message);
            onError();
        });
    };
    let apiServer = SERVER_API + "v1/upload";
</script>

<style scoped lang="scss">
    .upload {
        padding-top: $header-height;
        width: 90%;
        height: auto;

        margin-left: 5%;
        margin-right: 5%;
        max-height: 90%;
        overflow-y: scroll;

        .uploadedList {
            width: 90%;
            position: absolute;
            max-height: 60%;
            overflow-y: scroll;
        }
    }

</style>