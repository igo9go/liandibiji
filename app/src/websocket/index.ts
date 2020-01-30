import {Constants} from '../constants';
import {hideMessage, showMessage} from '../util/message';
import {destroyDialog, dialog} from '../util/dialog';
import {i18n} from '../i18n';
import {mountFile} from "../util/mount";

export class WebSocketUtil {
    public webSocket: WebSocket;
    private isFirst: boolean;

    constructor(liandi: ILiandi) {
        this.isFirst = true;
        this.connect(liandi);
    }

    private connect(liandi: ILiandi) {
        this.webSocket = new WebSocket(Constants.WEBSOCKET_ADDREDD);
        this.webSocket.onopen = () => {
            if (this.isFirst) {
                liandi.ws.webSocket.send(JSON.stringify({
                    cmd: 'dirs',
                    param: {},
                }));
            }
            this.isFirst = false;
        };
        this.webSocket.onclose = (e) => {
            console.warn('WebSocket is closed. Reconnect will be attempted in 1 second.', e);
            setTimeout(() => {
                this.connect(liandi);
            }, 1000);
        };
        this.webSocket.onerror = (err) => {
            console.error('WebSocket Error:', err);
            this.webSocket.close();
        };
        this.webSocket.onmessage = (event) => {
            const response = JSON.parse(event.data);
            if (response.code !== 0) {
                showMessage(response.msg, 0);
                return;
            }
            switch (response.cmd) {
                case 'mount':
                    liandi.navigation.onMount(liandi, response.data.url);
                    break;
                case 'mountremote':
                    hideMessage()
                    break;
                case 'ls':
                    liandi.files.onLs(liandi, response.data);
                    break;
                case 'get':
                    liandi.editors.onGet(liandi, response.data);
                    break;
                case 'dirs':
                    if (response.data.length === 0) {
                        dialog({
                            title: i18n[Constants.LANG].slogan,
                            content: `<div class="ft__center"><button class="button button--confirm">${i18n[Constants.LANG].mount}</button></div>`,
                            width: 400
                        });

                        const dialogElement = document.querySelector('#dialog');
                        dialogElement.querySelector('.button--confirm').addEventListener('click', () => {
                            mountFile(liandi.ws.webSocket);
                            destroyDialog();
                        });
                        return;
                    }
                    liandi.navigation.element.innerHTML = '';
                    response.data.forEach((url: string) => {
                        liandi.navigation.onMount(liandi, url);
                    });
                    break;
                case 'rename':
                    liandi.files.onRename(liandi, response.data);
                    break;
                case 'create':
                case 'remove':
                case 'mkdir':
                    window.liandi.liandi.ws.webSocket.send(JSON.stringify({
                        cmd: 'ls',
                        param: {
                            url: response.data.url,
                            path: response.data.path,
                        },
                    }));
                    break;
            }
        };
    }
}
