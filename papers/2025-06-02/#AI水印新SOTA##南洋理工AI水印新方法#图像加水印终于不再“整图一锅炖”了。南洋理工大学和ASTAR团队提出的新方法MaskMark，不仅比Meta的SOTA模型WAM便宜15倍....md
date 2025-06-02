# #AI水印新SOTA##南洋理工AI水印新方法#图像加水印终于不再“整图一锅炖”了。南洋理工大学和A*STAR团队提出的新方法MaskMark，不仅比Meta的SOTA模型WAM便宜15倍...

**URL**: https://weibo.com/6105753431/Pun0t5qPL

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E6%B0%B4%E5%8D%B0%E6%96%B0SOTA%23&amp;extparam=%23AI%E6%B0%B4%E5%8D%B0%E6%96%B0SOTA%23" data-hide=""><span class="surl-text">#AI水印新SOTA#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%8D%97%E6%B4%8B%E7%90%86%E5%B7%A5AI%E6%B0%B4%E5%8D%B0%E6%96%B0%E6%96%B9%E6%B3%95%23&amp;extparam=%23%E5%8D%97%E6%B4%8B%E7%90%86%E5%B7%A5AI%E6%B0%B4%E5%8D%B0%E6%96%B0%E6%96%B9%E6%B3%95%23" data-hide=""><span class="surl-text">#南洋理工AI水印新方法#</span></a><br><br>图像加水印终于不再“整图一锅炖”了。南洋理工大学和A*STAR团队提出的新方法MaskMark，不仅比Meta的SOTA模型WAM便宜15倍训练成本，性能还全面碾压。<br><br>该方法核心是“掩码机制”。训练时告诉模型“水印藏在哪”，模型学会精准嵌入和提取，像长了眼一样识别水印区域。它分两种版本：MaskMark-D全图嵌入、局部提取，MaskMark-ED只嵌入指定区域，更适合人脸、LOGO保护等。<br><br>性能方面，MaskMark在高保真度下依然维持接近100%提取准确率。局部水印场景中，尤其在小区域嵌入任务上，远超现有模型。水印定位（用IoU衡量）也稳压WAM和EditGuard，更神奇的是，它原本没为多水印设计，照样能一次性埋5个水印也不出错。<br><br>训练效率同样惊艳：单张A6000显卡20小时搞定，TFLOPs效率是WAM的15倍。支持不同比特长度、支持快速微调，对抗新攻击也能迅速适应。 <a href="https://weibo.com/ttarticle/p/show?id=2309405172382813257787" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">1/15成本，实现AI水印新SOTA | 南洋理工大学＆A*STAR</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3ly1i1yp1kwo9nj30jw0b7myl.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

南洋理工大学与A*STAR团队提出新型AI水印方法MaskMark，采用"掩码机制"精准定位水印区域，性能全面超越Meta的WAM模型。该方法分两种版本：MaskMark-D支持全图嵌入局部提取，MaskMark-ED专用于指定区域保护。实验显示其在高保真度下仍保持近100%提取准确率，尤其在小区域水印任务中表现突出，单卡训练仅需20小时，效率达WAM的15倍。此外，该方法意外展现出优秀的多水印处理能力，可同时嵌入5个水印且保持高定位精度（IoU指标领先），并具备快速适应新型攻击的微调能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-02T01:31:21Z
- **目录日期**: 2025-06-02
