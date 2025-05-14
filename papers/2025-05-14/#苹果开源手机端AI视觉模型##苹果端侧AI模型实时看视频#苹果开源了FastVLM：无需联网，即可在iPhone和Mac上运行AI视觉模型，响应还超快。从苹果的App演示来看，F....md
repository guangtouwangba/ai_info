# #苹果开源手机端AI视觉模型##苹果端侧AI模型实时看视频#苹果开源了FastVLM：无需联网，即可在iPhone和Mac上运行AI视觉模型，响应还超快。从苹果的App演示来看，F...

**URL**: https://weibo.com/6105753431/PrNs8tCMY

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%8B%B9%E6%9E%9C%E5%BC%80%E6%BA%90%E6%89%8B%E6%9C%BA%E7%AB%AFAI%E8%A7%86%E8%A7%89%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E8%8B%B9%E6%9E%9C%E5%BC%80%E6%BA%90%E6%89%8B%E6%9C%BA%E7%AB%AFAI%E8%A7%86%E8%A7%89%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#苹果开源手机端AI视觉模型#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%8B%B9%E6%9E%9C%E7%AB%AF%E4%BE%A7AI%E6%A8%A1%E5%9E%8B%E5%AE%9E%E6%97%B6%E7%9C%8B%E8%A7%86%E9%A2%91%23&amp;extparam=%23%E8%8B%B9%E6%9E%9C%E7%AB%AF%E4%BE%A7AI%E6%A8%A1%E5%9E%8B%E5%AE%9E%E6%97%B6%E7%9C%8B%E8%A7%86%E9%A2%91%23" data-hide=""><span class="surl-text">#苹果端侧AI模型实时看视频#</span></a><br><br>苹果开源了FastVLM：无需联网，即可在iPhone和Mac上运行AI视觉模型，响应还超快。<br><br>从苹果的App演示来看，FastVLM可以——<br><br>1、手势识别：对准用户的手部，模型可实时识别举起的手指数量。【图2】<br>2、图像文字识别：摄像头对准一张便签，FastVLM可在791毫秒内识别出文字信息。【图3】<br>3、图像描述生成：当模型看到一组emoji表情时，能即时生成英文简述，整个过程控制在800毫秒以内。【图4】<br><br>FastVLM的主要技术优势包括：<br><br>- 视觉token优化：通过减少输出token数量，最多可压缩至原来的1/16，降低了大模型在图像预处理环节的资源消耗。<br>    <br>- 推理速度提升：在关键指标Time-to-First-Token（首次输出延迟）上，FastVLM比LLaVA-OneVision快85倍，响应更即时。<br>    <br>- 模型体积精简：最小规格的FastVLM视觉编码器相比传统模型缩小了3.4倍，而且性能未缩水。<br>    <br>- 原生本地推理支持：模型专为Apple Silicon芯片（如M1、M2）优化，兼容苹果机器学习框架。<br><br>架构方面，FastViTHD采用混合结构，结合了卷积和Transformer，配合多尺度下采样，既保留高分辨率信息，又压缩了冗余token。<br><br>在多个主流多模态评测基准（如GQA、TextVQA、SeedBench）中，FastVLM全面超越了包括ConvLLaVA、Cambrian-1、MM1在内的多款大模型，且资源消耗更低，推理效率更高。<br><br>目前，苹果在开源平台提供了三种模型规格（0.5B / 1.5B / 7B），同时发布了适配Apple Silicon的版本与iOS demo应用，方便开发者快速测试。<br><br>代码：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fapple%2Fml-fastvlm" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>论文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2412.13303" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1f4o448e6j31ba10kh2d.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1f4ocfcsvg30d40qwu0z.gif" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1f4od93pdg30d40qw4qs.gif" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1f4odte8mg30d40qwkjl.gif" referrerpolicy="no-referrer"><br><br>

## AI 摘要

苹果开源了手机端AI视觉模型FastVLM，可在iPhone和Mac本地运行，无需联网。该模型支持实时手势识别（如识别手指数量）、图像文字识别（791毫秒内完成）和图像描述生成（800毫秒内响应）。技术亮点包括：视觉token优化（压缩至1/16）、推理速度比LLaVA-OneVision快85倍、模型体积缩小3.4倍，并针对Apple Silicon芯片优化。采用混合架构（卷积+Transformer），在多项基准测试中超越同类模型。提供0.5B/1.5B/7B三种规格，已发布iOS演示应用和开源代码。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-14T13:12:49Z
- **目录日期**: 2025-05-14
