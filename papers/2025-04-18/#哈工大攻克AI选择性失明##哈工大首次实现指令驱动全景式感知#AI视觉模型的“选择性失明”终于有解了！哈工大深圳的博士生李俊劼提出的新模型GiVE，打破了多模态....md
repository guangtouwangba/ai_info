# #哈工大攻克AI选择性失明##哈工大首次实现指令驱动全景式感知#AI视觉模型的“选择性失明”终于有解了！哈工大深圳的博士生李俊劼提出的新模型GiVE，打破了多模态...

**URL**: https://weibo.com/6105753431/PnQLrc0RW

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%93%88%E5%B7%A5%E5%A4%A7%E6%94%BB%E5%85%8BAI%E9%80%89%E6%8B%A9%E6%80%A7%E5%A4%B1%E6%98%8E%23&amp;extparam=%23%E5%93%88%E5%B7%A5%E5%A4%A7%E6%94%BB%E5%85%8BAI%E9%80%89%E6%8B%A9%E6%80%A7%E5%A4%B1%E6%98%8E%23" data-hide=""><span class="surl-text">#哈工大攻克AI选择性失明#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%93%88%E5%B7%A5%E5%A4%A7%E9%A6%96%E6%AC%A1%E5%AE%9E%E7%8E%B0%E6%8C%87%E4%BB%A4%E9%A9%B1%E5%8A%A8%E5%85%A8%E6%99%AF%E5%BC%8F%E6%84%9F%E7%9F%A5%23&amp;extparam=%23%E5%93%88%E5%B7%A5%E5%A4%A7%E9%A6%96%E6%AC%A1%E5%AE%9E%E7%8E%B0%E6%8C%87%E4%BB%A4%E9%A9%B1%E5%8A%A8%E5%85%A8%E6%99%AF%E5%BC%8F%E6%84%9F%E7%9F%A5%23" data-hide=""><span class="surl-text">#哈工大首次实现指令驱动全景式感知#</span></a><br><br>AI视觉模型的“选择性失明”终于有解了！<br><br>哈工大深圳的博士生李俊劼提出的新模型GiVE，打破了多模态大模型只能关注图片中显著元素的局限，让AI真正“看清楚”图像的每个角落。<br><br>核心突破在于引入AG-Adapter模块，为视觉编码器加上一双“耳朵”，从此便能听懂文本指令，根据用户需求聚焦图像中容易被忽略的部分，比如角落里的物体或背景中的标识。<br><br>GiVE还设计了三个针对细节感知的训练目标（Loss函数）：<br><br>- 图像与文本之间的细节对齐；<br>- 图像内部多个目标间的关系提取；<br>- 对潜在目标的二分类识别。<br><br>为训练模型，团队打造了包含8万图像和24万标注的多目标指令数据集MOInst，覆盖264类物体。<br><br>更妙的是，GiVE可以作为插件集成到现有视觉系统中，简单灵活，不需要从零开始重训模型。<br><br>实验显示，GiVE显著提升了图文检索和图像分类任务的准确性，为AI从“粗略看图”进化到“按需识图”奠定了基础。<br><br>未来，它有望在医疗、安防、自动驾驶等场景大显身手，成为AI迈向通用智能的关键一环。<br><br>感兴趣的小伙伴可以点击——<br>代码：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2FAlephZr%2FGiVE%2Ftree%2Fmain" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>数据集：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fhuggingface.co%2Fdatasets%2FDF1024%2FMOInst" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0l3mzu1k0j31r00e8h3j.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0l3n16xxzj30zk0sw161.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

哈工大深圳博士生李俊劼团队提出GiVE模型，解决了AI视觉模型的"选择性失明"问题。该模型通过AG-Adapter模块使AI能根据文本指令聚焦图像细节（如角落物体），并设计了三个训练目标优化细节感知。团队构建了包含8万图像、24万标注的MOInst数据集。GiVE可作为插件集成到现有系统，显著提升了图文检索和图像分类准确率。这项突破为医疗、安防等领域的精细化视觉应用奠定了基础，相关代码和数据集已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-18T23:03:23Z
- **目录日期**: 2025-04-18
