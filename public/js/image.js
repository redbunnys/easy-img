const { createApp } = Vue

createApp({
    data() {
        return {
            isDragging: false,
            previews: [],
            history: [],
            previewId: 0
        }
    },
    mounted() {
        // 从 localStorage 加载历史记录
        const savedHistory = localStorage.getItem('uploadHistory')
        if (savedHistory) {
            this.history = JSON.parse(savedHistory)
        }

        // 添加剪贴板粘贴事件监听
        document.addEventListener('paste', this.handlePaste)
    },
    unmounted() {
        // 组件销毁时移除事件监听
        document.removeEventListener('paste', this.handlePaste)
    },
    methods: {
        // 添加处理粘贴事件的方法
        handlePaste(e) {
            const items = e.clipboardData.items
            
            for (let item of items) {
                // 检查是否是图片
                if (item.type.indexOf('image') !== -1) {
                    const file = item.getAsFile()
                    this.uploadFile(file)
                    
                    // 阻止默认粘贴行为
                    e.preventDefault()
                    break
                }
            }
        },
        triggerFileInput() {
            // 先清空 input 的值，确保同一文件可以重复选择
            this.$refs.fileInput.value = ''
            this.$refs.fileInput.click()
        },
        handleDrop(e) {
            this.isDragging = false
            const files = [...e.dataTransfer.files]
            files.forEach(this.uploadFile)
        },
        handleFiles(e) {
            const files = [...e.target.files]
            if (files.length > 0) {
                files.forEach(this.uploadFile)
            }
        },
        async uploadFile(file) {
            // 验证文件类型和大小
            if (!file.type.match('image.*')) {
                this.showToast('请上传图片文件！')
                return
            }
            if (file.size > 5 * 1024 * 1024) {
                this.showToast('文件大小不能超过5MB！')
                return
            }

            // 清除之前的预览
            this.previews.forEach(p => URL.revokeObjectURL(p.url))
            this.previews = []

            // 创建新的预览
            const previewId = this.previewId++
            const preview = {
                id: previewId,
                url: URL.createObjectURL(file),
                status: 'uploading'
            }
            this.previews.push(preview)

            // 上传文件
            const formData = new FormData()
            formData.append('image', file)

            try {
                const response = await fetch('/api/upload', {
                    method: 'POST',
                    body: formData
                })

                if (!response.ok) {
                    throw new Error('Upload failed')
                }

                const result = await response.json()
                
                // 更新预览状态
                preview.status = 'success'
                preview.fileUrl = result.url
                
                // 添加到历史记录
                const historyItem = {
                    id: Date.now(),
                    url: result.url
                }
                this.history.unshift(historyItem)
                
                // 保存到 localStorage
                localStorage.setItem('uploadHistory', JSON.stringify(this.history.slice(0, 20)))
                
                this.showToast('上传成功！')

            } catch (error) {
                console.error('Upload error:', error)
                preview.status = 'error'
                this.showToast('上传失败，请重试')
            }
        },
        removePreview(id) {
            const index = this.previews.findIndex(p => p.id === id)
            if (index !== -1) {
                URL.revokeObjectURL(this.previews[index].url)
                this.previews.splice(index, 1)
            }
        },
        async copyToClipboard(text, format = 'url') {
            try {
                const url = new URL(text, window.location.origin).toString();
                let copyText;
                
                switch (format) {
                    case 'markdown':
                        copyText = `![image](${url})`;
                        break;
                    case 'html':
                        copyText = `<img src="${url}" alt="image"/>`;
                        break;
                    default:
                        copyText = url;
                }
                
                await navigator.clipboard.writeText(copyText);
                this.showToast(`已复制${format.toUpperCase()}格式到剪贴板`);
            } catch (err) {
                console.error('Copy failed:', err);
                this.showToast('复制失败，请手动复制');
            }
        },
        showToast(message) {
            const toast = document.createElement('div')
            toast.className = 'toast'
            toast.textContent = message
            document.body.appendChild(toast)

            setTimeout(() => toast.classList.add('show'), 10)
            setTimeout(() => {
                toast.classList.remove('show')
                setTimeout(() => toast.remove(), 300)
            }, 3000)
        },
        clearHistory() {
            this.history = []
            localStorage.removeItem('uploadHistory')
            this.showToast('历史记录已清除')
        },
        deleteHistoryItem(id) {
            const index = this.history.findIndex(item => item.id === id)
            if (index !== -1) {
                this.history.splice(index, 1)
                // 更新 localStorage
                localStorage.setItem('uploadHistory', JSON.stringify(this.history))
                this.showToast('已删除该记录')
            }
        }
    }
}).mount('#app')
