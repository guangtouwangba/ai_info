# Infinity: Scaling Bitwise AutoRegressive Modeling for High-Resolution Image Synthesis网页链接本文介绍了一种名为Infinity的位运算视觉自回归模型，该模型...

**URL**: https://weibo.com/1870858943/P4uAXfjkX

## 原始摘要

Infinity: Scaling Bitwise AutoRegressive Modeling for High-Resolution Image Synthesis<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fwww.aminer.cn%2Fpub%2F67526da4ae8580e7ff3d4f6e%2F%3Ff%3Dwb" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>本文介绍了一种名为Infinity的位运算视觉自回归模型，该模型能够根据语言指令生成高分辨率、逼真的图像。Infinity在位运算标记预测框架下重新定义了视觉自回归模型，通过无限词汇量的标记分类器和位运算自我校正机制，显著提高了生成能力和细节表现。该方法理论上将标记词汇量扩展至无限，并同步扩大了变换器规模，相较于传统的自回归模型，展现出了更强大的扩展能力。Infinity在自回归文本到图像模型中创下新纪录，性能超过了顶尖的扩散模型如SD3-Medium和SDXL。特别地，Infinity将GenEval基准分数从0.62提升至0.73，ImageReward基准分数从0.87提升至0.96，在0.8秒内完成1024x1024图像的生成，速度是SD3-Medium的2.6倍，成为最快的文本到图像模型。研究团队将发布模型和代码，以推动Infinity在视觉生成和统一标记建模领域的进一步探索。<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%9B%9E%E5%BD%92%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#回归模型#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%89%A9%E6%95%A3%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E6%89%A9%E6%95%A3%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#扩散模型#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%AF%AD%E8%A8%80%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#语言模型#</span></a><a href="https://m.weibo.cn/p/index?extparam=%E4%BA%BA%E5%B7%A5%E6%99%BA%E8%83%BD&amp;containerid=100808f068f0dad74789bee210163c40a4b50d" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">人工智能</span></a><a href="https://m.weibo.cn/p/index?extparam=%E7%A1%95%E5%A3%AB%E8%AE%BA%E6%96%87&amp;containerid=1008084cacf38f5903dc7b04550404d0bd3608" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">硕士论文</span></a><a href="https://m.weibo.cn/p/index?extparam=%E8%AE%BA%E6%96%87%E5%86%99%E4%BD%9C&amp;containerid=1008084f70c9f305ba97c50dbca8c25c8747d7" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">论文写作</span></a><img style="" src="https://tvax4.sinaimg.cn/large/6f830abfly1hwhxxq7oocj22cv18f1ky.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

本文介绍了一种名为Infinity的位运算视觉自回归模型，该模型通过无限词汇量的标记分类器和位运算自我校正机制，显著提高了生成高分辨率、逼真图像的能力。Infinity在自回归文本到图像模型中创下新纪录，性能超过了顶尖的扩散模型如SD3-Medium和SDXL，特别是在GenEval和ImageReward基准测试中表现优异。此外，Infinity在0.8秒内完成1024x1024图像的生成，速度是SD3-Medium的2.6倍，成为最快的文本到图像模型。研究团队将发布模型和代码，以推动Infinity在视觉生成和统一标记建模领域的进一步探索。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-21T01:32:41Z
- **目录日期**: 2025-03-21
