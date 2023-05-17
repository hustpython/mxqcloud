<template>
    <Header/>
    <div class="upload">
        <n-upload
                style="padding-top: 10px;max-height: 300px;overflow-y: scroll"
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
            <n-collapse arrow-placement="right" default-expanded-names="1">
                <n-collapse-item :title="completdCount" name="1">

                    <n-space justify="end">
                        <n-button v-show="isSelected && uploadCompleteData.length>0" size="tiny"
                                  @click=handleDelete>
                            删除
                        </n-button>
                        <n-checkbox size="small"
                                    v-show="uploadCompleteData.length>0"
                                    v-model:checked="allCheck"
                                    @update:checked="handleCheckedChange">
                            全选
                        </n-checkbox>
                    </n-space>
                    <n-divider/>
                    <div completeDatail v-for="(item) in uploadCompleteData">
                        <n-space justify="space-between">
                            <!--                            <n-image-->
                            <!--                                    width=51-->
                            <!--                                    :src="getFileUrl(item.Url)"/>-->
                            <n-space vertical>
                                <n-ellipsis style="max-width: 180px;font-size: 12px">
                                    {{ item.FileName }}
                                </n-ellipsis>
                                <div style="font-size: 12px;color: gray">
                                    {{item.DateStr}} {{item.SizeStr}}
                                </div>
                            </n-space>
                            <n-checkbox v-model:checked="item.Check" size="small"
                                        @update:checked="handleSingleCheckedChange"/>
                        </n-space>
                    </div>

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
    import {onBeforeMount, computed, ref} from 'vue'
    import {GetDevice} from '@/util/Device.js'
    import {GetUploadData, DelUploadData} from '@/util/api.js'

    let res = GetDevice();


    const completdCount = computed(() => {
        return "已完成 " + uploadCompleteData.value.length
    });
    const uploadCompleteData = ref([]);
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
            onFinish();
            uploadCompleteData.value.push(json);
        }).catch((error) => {
            onError();
        });
    };
    let apiServer = SERVER_API + "v1/upload";

    const getFileUrl = (url) => {
        return SERVER_API + url
    }

    onBeforeMount(() => {
        GetUploadData(res).then((ress) => {
            if (ress.code === 200) {
                uploadCompleteData.value = ress.data;
            }
        });
    });

    const isSelected = ref(false);
    const allCheck = ref(false);

    const handleCheckedChange = (value) => {
        if (value) {
            uploadCompleteData.value.forEach((k, v) => {
                uploadCompleteData.value[v].Check = true
            })
            isSelected.value = true
        } else {
            uploadCompleteData.value.forEach((k, v) => {
                uploadCompleteData.value[v].Check = false
            })
            isSelected.value = false;
        }

    }

    const handleSingleCheckedChange = (value) => {
        if (!value) {
            allCheck.value = false;
        } else {
            isSelected.value = true;
        }
        let noCheck = true
        uploadCompleteData.value.forEach((k) => {
            if (k.Check) {
                noCheck = false
            }
        })
        if (noCheck) {
            isSelected.value = false;
        }
    }

    const handleDelete = async () => {
        for (let i = uploadCompleteData.value.length - 1; i >= 0; i--) {
            if (uploadCompleteData.value[i].Check) {
                await DelUploadData(uploadCompleteData.value[i].Id).then((ress) => {
                    if (ress.code === 200) {
                        uploadCompleteData.value.splice(i, 1)
                    }
                });
            }
        }
    }
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
            max-height: 50%;
            overflow-y: scroll;
        }
    }

</style>