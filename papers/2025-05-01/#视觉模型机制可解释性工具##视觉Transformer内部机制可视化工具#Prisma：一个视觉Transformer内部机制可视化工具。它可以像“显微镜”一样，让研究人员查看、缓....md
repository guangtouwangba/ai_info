# #视觉模型机制可解释性工具##视觉Transformer内部机制可视化工具#Prisma：一个视觉Transformer内部机制可视化工具。它可以像“显微镜”一样，让研究人员查看、缓...

**URL**: https://weibo.com/6105753431/PpFgdcPhL

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%A7%86%E8%A7%89%E6%A8%A1%E5%9E%8B%E6%9C%BA%E5%88%B6%E5%8F%AF%E8%A7%A3%E9%87%8A%E6%80%A7%E5%B7%A5%E5%85%B7%23&amp;extparam=%23%E8%A7%86%E8%A7%89%E6%A8%A1%E5%9E%8B%E6%9C%BA%E5%88%B6%E5%8F%AF%E8%A7%A3%E9%87%8A%E6%80%A7%E5%B7%A5%E5%85%B7%23" data-hide=""><span class="surl-text">#视觉模型机制可解释性工具#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%A7%86%E8%A7%89Transformer%E5%86%85%E9%83%A8%E6%9C%BA%E5%88%B6%E5%8F%AF%E8%A7%86%E5%8C%96%E5%B7%A5%E5%85%B7%23&amp;extparam=%23%E8%A7%86%E8%A7%89Transformer%E5%86%85%E9%83%A8%E6%9C%BA%E5%88%B6%E5%8F%AF%E8%A7%86%E5%8C%96%E5%B7%A5%E5%85%B7%23" data-hide=""><span class="surl-text">#视觉Transformer内部机制可视化工具#</span></a><br><br>Prisma：一个视觉Transformer内部机制可视化工具。<br><br>它可以像“显微镜”一样，让研究人员查看、缓存和干预视觉模型的内部激活，从而逆向分析模型学到的算法。<br><br>Prisma的灵感来源于TransformerLens。不过Prisma更专注于视觉、视频模型，而TransformerLens则专注于语言模型。<br><br>该工具具有以下特点：<br><br>- 一整套分析工具：它不仅有可视化功能（比如注意力热图、特征重建图），还提供了各种机制层面分析工具，像activation patching、logit lens、circuit tracing、画出SAE重建前后的输入对比等。<br><br>- 支持视频Transformer：Prisma支持ViViT和V-JEPA等视频模型，意味着你可以在时间维度上观察特征如何变化，理解模型如何捕捉动态信息。<br><br>- 内置“小模型”方便调试：类似语言圈常用的“用小模型做机制探索”，Prisma也提供了1–4层的toy ViTs，方便快速验证思路、寻找表示“回路”（circuits）。<br><br>研究人员在应用Prisma后，发现了以下几个有趣的事：<br><br>1. 视觉模型比语言模型“更活跃”：Prisma训练的Sparse Autoencoder（SAE）揭示，像CLIP-B/32这类模型，每个patch活跃的latent可以超过500个，而语言模型的token大多只有十几个。这说明视觉表征在分布上更加密集。<br><br>2. SAE不仅能还原特征，还能降loss：传统上我们以为，插入SAE重建的表示可能会破坏模型性能，结果Prisma发现，这反而能显著降低cross-entropy loss。这提示我们，特征重建也可能起到某种“降噪”效果。<br><br>3. CLS token和Patch token的行为完全不同 ：在ViT中，随着层数增加，CLS token会越来越“饱和”、吸收大量特征表示，而普通patch token则逐渐“沉寂”。这与语言模型中token，相对均衡的特征分布形成鲜明对比，揭示了视觉Transformer中独特的信息聚合机制。<br><br>对于关注模型内部表示、想搞清楚“模型到底在干什么”的研究者来说，Prisma不仅复制了NLP interpretability的成功路径，也为CV模型的“脑科学”研究打开了新局面。<br><br>该项目完全开源，感兴趣小伙伴可以点击——<br>Prisma：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fprismaproject%2Fprisma" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>论文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2504.19475" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>TransformerLens：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2FTransformerLensOrg%2FTransformerLens" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0yx2bbxdmj30zk0i20zb.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Prisma是一个视觉Transformer模型内部机制可视化工具，可像"显微镜"一样观察、干预和分析视觉模型的内部激活。它支持视频模型（如ViViT/V-JEPA），提供注意力热图、特征重建等分析工具，并内置小型ViT模型用于调试。研究发现：1）视觉模型比语言模型激活更密集；2）特征重建可能提升性能；3）ViT中CLS token与patch token行为差异显著。该工具为CV模型可解释性研究开辟了新途径，已开源发布。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T01:31:11Z
- **目录日期**: 2025-05-01
