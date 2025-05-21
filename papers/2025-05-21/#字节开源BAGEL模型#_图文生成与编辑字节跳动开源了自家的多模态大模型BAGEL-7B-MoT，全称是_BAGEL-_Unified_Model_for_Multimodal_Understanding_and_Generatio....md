# #字节开源BAGEL模型# 图文生成与编辑字节跳动开源了自家的多模态大模型BAGEL-7B-MoT，全称是 BAGEL: Unified Model for Multimodal Understanding and Generatio...

**URL**: https://weibo.com/6105753431/PsQmMfysW

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%AD%97%E8%8A%82%E5%BC%80%E6%BA%90BAGEL%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E5%AD%97%E8%8A%82%E5%BC%80%E6%BA%90BAGEL%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#字节开源BAGEL模型#</span></a> 图文生成与编辑<br><br>字节跳动开源了自家的多模态大模型BAGEL-7B-MoT，全称是 BAGEL: Unified Model for Multimodal Understanding and Generation，定位是图文生成与编辑模型。<br><br>- 多模态能力：BAGEL在视觉问答、图像理解等基准测试中超越了Qwen2.5-VL和InternVL-2.5，在图生图质量上也能和SD3 一战；<br><br>- 支持自由图像编辑：除了基础的文生图，还能做智能修改、视角变换，甚至可以处理连续画面推理和“世界建模”等任务；<br><br>- 底层结构是MoT（Mixture-of-Transformer-Experts），即专家混合架构 + 双编码器设计（分别处理像素级和语义级特征）；<br><br>- 训练数据覆盖语言、图像、视频、网页等多模态大规模混合内容，总参数14B，活跃参数7B；<br><br>值得注意的是，BAGEL在图像编辑方面表现亮眼（GEdit-Bench 与IntelligentBench都接近或优于Gemini），并且在训练过程中出现了“能力涌现”现象：从基础生成理解，到智能编辑，功能随着训练阶段层层展开。<br><blockquote> - 转发 <a href="https://weibo.com/2169039837" target="_blank">@karminski-牙医</a>: 开始扎堆了啊，字节跳动12小时前刚发了个新模型——BAGEL-7B-MoT，给大家带来实测！<br><br>这是个混合专家多模态模型，支持视觉理解，文本到图像生成，图像编辑，并且思考模式可以选择开启。官方说要比  Qwen2.5-VL 和 InternVL-2.5 表现好。这个模型本身是基于 Qwen2.5-7B-Instruct 和 siglip-so400m-14-980-flash-attn2-navit 模型微调的，并使用 FLUX.1-schnell VAE 模型。<br><br>文本生成 图1，问了模型名称合知识截止日期。这个没什么好说的，毕竟是个多模态模型<br><br>图片生成 图2，生成一个热气球吊着一个房子在天空中飞行的图片，指令遵循可以，没出现奇怪的东西<br><br>图片理解 图3，随手传了个防水马粪袋，问这是什么，推理结果相当不错，而且输出很精炼，不至于说一大堆没用的东西，这个非常适合嵌到工作流中。<br><br>图片修改 图4，图5，给三棱镜增加一束从底部的入射光。这个就差一些了。我先要求了增加入射光，这个指令遵循很严格，无任何发挥，因此没渲染色散效果。而且入射光貌似方向反了？然后又要求增加色散效果，不过还是没有正确渲染。<br><br>总结：我这个小测试中来看作为图片推理应该是最强的，其余次之，图片修改不太满意。考虑到模型只有7B，感兴趣的同学可以试试。如果量化后装到16G的显卡里还是很轻松的。<br><br>论文地址：huggingface.co/papers/2505.14683<br>模型地址：huggingface.co/ByteDance-Seed/BAGEL-7B-MoT<br>repo: github.com/bytedance-seed/BAGEL<br><br><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23ai%E5%88%9B%E9%80%A0%E8%90%A5%23" data-hide=""><span class="surl-text">#ai创造营#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E7%94%9F%E6%B4%BB%E6%8C%87%E5%8D%97%23&amp;extparam=%23AI%E7%94%9F%E6%B4%BB%E6%8C%87%E5%8D%97%23" data-hide=""><span class="surl-text">#AI生活指南#</span></a><img style="" src="https://tvax3.sinaimg.cn/large/8148ebddgy1i1mzu1077tj21jv0d7aix.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/8148ebddgy1i1mzu1p0eij21jn1ikqv5.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/8148ebddgy1i1mzu1s079j21jo1ktqv5.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/8148ebddgy1i1mzu1e4xij21jv1rb1kx.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/8148ebddgy1i1mzu1m6bvj21k01owx13.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/8148ebddgy1i1mzu24883j20wk1u31ky.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/8148ebddgy1i1mzu11qkcj20we1cak2y.jpg" referrerpolicy="no-referrer"><br><br></blockquote>

## AI 摘要

字节跳动开源了多模态大模型BAGEL-7B-MoT，具备图文生成与编辑能力。该模型在视觉问答和图像理解任务上超越Qwen2.5-VL等竞品，支持文生图、智能修改、视角变换等编辑功能。采用专家混合架构（MoT）和双编码器设计，训练数据涵盖多模态内容，总参数14B（活跃7B）。实测显示其在图像理解表现突出，但编辑能力仍有改进空间。模型基于Qwen2.5-7B微调，量化后可适配16G显存设备。论文和代码已公开在Hugging Face和GitHub。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T07:02:51Z
- **目录日期**: 2025-05-21
