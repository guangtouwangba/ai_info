# #一秒转写六十分钟录音##英伟达开源语音识别模型登顶OpenASR#短短一秒钟，转写六十分钟录音？！英伟达开源6亿参数的语音识别模型Parakeet TDT 0.6B V2，模型在Hu...

**URL**: https://weibo.com/6105753431/PqRHsbaOi

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%80%E7%A7%92%E8%BD%AC%E5%86%99%E5%85%AD%E5%8D%81%E5%88%86%E9%92%9F%E5%BD%95%E9%9F%B3%23&amp;extparam=%23%E4%B8%80%E7%A7%92%E8%BD%AC%E5%86%99%E5%85%AD%E5%8D%81%E5%88%86%E9%92%9F%E5%BD%95%E9%9F%B3%23" data-hide=""><span class="surl-text">#一秒转写六十分钟录音#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%8B%B1%E4%BC%9F%E8%BE%BE%E5%BC%80%E6%BA%90%E8%AF%AD%E9%9F%B3%E8%AF%86%E5%88%AB%E6%A8%A1%E5%9E%8B%E7%99%BB%E9%A1%B6OpenASR%23&amp;extparam=%23%E8%8B%B1%E4%BC%9F%E8%BE%BE%E5%BC%80%E6%BA%90%E8%AF%AD%E9%9F%B3%E8%AF%86%E5%88%AB%E6%A8%A1%E5%9E%8B%E7%99%BB%E9%A1%B6OpenASR%23" data-hide=""><span class="surl-text">#英伟达开源语音识别模型登顶OpenASR#</span></a><br><br>短短一秒钟，转写六十分钟录音？！<br><br>英伟达开源6亿参数的语音识别模型Parakeet TDT 0.6B V2，模型在HuggingFace Open-ASR榜单上以RTFx 3380的成绩登顶！<br><br>这款语音模型适用于多种需要语音转文字功能的场景，包括但不限于语音助手、会议记录、采访整理、字幕生成和客服质检……<br><br>它具有三个特点：<br>1、字级时间戳精准预测：每个单词的出现时间都能精确定位<br>2、智能标点与自动大写：口语转文字自动加句号逗号，自动首字母大写<br>3、超强特殊场景识别：口语数字准确转换、歌词转录都不在话下<br><br>画重点：该模型基于CC-BY-4.0协议开源，可以用于商业用途！<br><br>再来看一些技术层面的亮点：<br>- 采用FastConformer编码器+TDT解码器架构，可处理长达24分钟的音频片段<br>- 总训练数据量达12万小时，结合了人工标注和伪标注数据源，包括LibriSpeech、Fisher、YTC、YODAS等数据集。<br>- 可通过NVIDIA NeMo获取，针对GPU推理进行了优化，使用pip install -U nemo_toolkit['asr']命令即可安装。<br>- 兼容Linux系统，支持Ampere、Blackwell、Hopper、Volta GPU架构，最低需要2GB显存。<br><br>用于训练的Granary数据集还会在Interspeech 2025会议后公开。<br><br>需要提醒一下各位，想要训练、微调或体验这个模型，需要先安装NVIDIA NeMo框架。建议在安装最新版PyTorch后再进行安装哦~<br><br>Demo链接：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fhuggingface.co%2Fspaces%2Fnvidia%2Fparakeet-tdt-0.6b-v2" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>项目主页：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fhuggingface.co%2Fnvidia%2Fparakeet-tdt-0.6b-v2" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i181pjheytj30ok0zkwq0.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i181pl5xrxj30t10zkam4.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

英伟达开源了6亿参数的语音识别模型Parakeet TDT 0.6B V2，在Open-ASR榜单以RTFx 3380登顶。该模型支持1秒转写60分钟录音，具备字级时间戳、智能标点、特殊场景识别（如数字转换）等功能，适用于语音助手、会议记录等场景。基于CC-BY-4.0协议开源，可商用。技术亮点包括FastConformer编码器+TDT解码器架构，支持24分钟长音频处理，训练数据达12万小时。需通过NVIDIA NeMo框架安装，兼容Linux系统，最低需2GB显存。Demo和项目主页已开放。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-08T08:03:05Z
- **目录日期**: 2025-05-08
