"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
document.getElementById('uploadBtn').addEventListener('click', () => __awaiter(void 0, void 0, void 0, function* () {
    var _a, _b;
    const fileInput = document.getElementById('fileInput');
    const file = (_b = (_a = fileInput.files) === null || _a === void 0 ? void 0 : _a[0]) !== null && _b !== void 0 ? _b : null;
    if (!file) {
        alert('No file selected');
        return;
    }
    const totalsize = file.size;
    let uploaded = 0;
    const stream = file.stream();
    const reader = stream.getReader();
    const uploadStream = new ReadableStream({
        pull(controller) {
            return __awaiter(this, void 0, void 0, function* () {
                const { done, value } = yield reader.read();
                if (done) {
                    controller.close();
                }
                if (value) {
                    uploaded += value.length;
                    updateProgress(uploaded, totalsize);
                }
            });
        }
    });
    console.log(`hi: ${file.name}`);
    yield fetch(`/upload_stream?file_name=${encodeURIComponent(file.name)}`, {
        method: 'POST',
        body: uploadStream,
        duplex: 'half'
    });
    alert('Upload completed');
}));
function updateProgress(uploaded, totalsize) {
    const percent = (uploaded / totalsize) * 100;
    const progressBar = document.getElementById('progressBar');
    progressBar.style.width = `${percent}%`;
}
