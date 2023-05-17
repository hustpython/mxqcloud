<template>
    <Header/>
    <div class="download">
        <n-collapse arrow-placement="right" :default-expanded-names=expandNames>
            <n-collapse-item :title="item.Name" :name="item.Name" v-for="(item,index) in downloadList">
                <n-space justify="end">
                    <n-button v-show="item.IsSelected && downloadList.length>0" size="tiny"
                              @click=handleDownload(index)>
                        下载
                    </n-button>
                    <n-checkbox size="small"
                                v-show="downloadList.length>0"
                                v-model:checked="item.AllCheck"
                                @update:checked="handleCheckedChange(index)">
                        全选
                    </n-checkbox>
                </n-space>
                <n-divider/>
                <div completeDatail v-for="(item1,index1) in item.Uploads">
                    <n-space justify="space-between">
                        <!--                        <n-image-->
                        <!--                                width=51-->
                        <!--                                :src="getFileUrl(item1.Url)"/>-->
                        <n-space vertical>
                            <n-ellipsis style="max-width: 180px;font-size: 12px">
                                {{ item1.FileName }}
                            </n-ellipsis>
                            <div style="font-size: 12px;color: gray">
                                {{item1.DateStr}} {{item1.SizeStr}}
                            </div>
                            <n-progress v-show="item1.Progress!==undefined" type="line" :percentage="item1.Progress">
                            </n-progress>
                        </n-space>
                        <n-space justify="end">
                            <n-collapse>
                                <n-collapse-item title="预览">
                                    <n-image
                                            width=51
                                            :src="getFileUrl(item1.Url)"/>
                                </n-collapse-item>
                            </n-collapse>

                            <n-checkbox v-model:checked="item1.Check" size="small"
                                        @update:checked="handleSingleCheckedChange(index,index1)"/>
                        </n-space>
                    </n-space>
                </div>
            </n-collapse-item>
        </n-collapse>
    </div>
    <Footer/>
</template>

<script setup>
    import Header from '@/components/Header.vue'
    import Footer from '@/components/Footer.vue'
    import SERVER_API from '@/util/config.js'
    import {GetAllDownload} from '@/util/api.js'
    import {onBeforeMount, ref} from 'vue'
    import {saveAs} from 'file-saver';
    import {lyla} from "lyla";


    const downloadList = ref([]);
    const expandNames = [];
    onBeforeMount(() => {
        GetAllDownload().then((ress) => {
            if (ress.code === 200) {
                downloadList.value = ress.data;
                downloadList.value.forEach((k, v) => {
                    downloadList.value[v].AllCheck = false;
                    downloadList.value[v].IsSelected = false;
                    expandNames.push(k.Name);
                })
            }
        });
    });

    const getFileUrl = (url) => {
        return SERVER_API + url
    }

    const handleCheckedChange = (index) => {
        if (downloadList.value[index].AllCheck) {
            downloadList.value[index].Uploads.forEach((k, v) => {
                downloadList.value[index].Uploads[v].Check = true;
            })
            downloadList.value[index].IsSelected = true;
        } else {
            downloadList.value[index].Uploads.forEach((k, v) => {
                downloadList.value[index].Uploads[v].Check = false
            })
            downloadList.value[index].IsSelected = false;
        }
    }

    const handleSingleCheckedChange = (i1, i2) => {
        if (!downloadList.value[i1].Uploads[i2].Check) {
            downloadList.value[i1].AllCheck = false;
        } else {
            downloadList.value[i1].IsSelected = true;
        }

        let noCheck = true
        downloadList.value[i1].Uploads.forEach((k) => {
            if (k.Check) {
                noCheck = false
            }
        })
        if (noCheck) {
            downloadList.value[i1].IsSelected = false;
        }
    }

    const handleDownload = (index) => {
        downloadList.value[index].Uploads.forEach((k, v) => {
            if (k.Check) {
                //saveAs(, k.FileName);
                downloadProgress(getFileUrl(k.Url), k.FileName, index, v);
            }
        })

    }

    const downloadProgress = async (url, fileName, i1, i2) => {
        await lyla.get(url, {
            responseType: "blob",
            onDownloadProgress: ({percent}) => {
                console.log(percent, fileName, i1, i2)
                downloadList.value[i1].Uploads[i2].Progress = Math.ceil(percent);
            }
        }).then((res) => {
            let blob = new Blob([res.body]);
            saveAs(blob, fileName)
        }).catch((error) => {

        });
    }
</script>

<style scoped lang="scss">
    .download {
        padding-top: calc($header-height + 15px);
        width: 90%;
        height: auto;
        margin-left: 5%;
        margin-right: 5%;
        position: absolute;
        max-height: 90%;
        overflow-y: scroll
    }
</style>