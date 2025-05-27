# #首个多模态扩散大语言模型##文本生成新范式#文本生成新范式——扩散大语言模型（dLLM），你听过吗？普林斯顿大学、字节Seed、北大、清华等团队，搞出来个多模态...

**URL**: https://weibo.com/6105753431/PtLMCFApE

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E9%A6%96%E4%B8%AA%E5%A4%9A%E6%A8%A1%E6%80%81%E6%89%A9%E6%95%A3%E5%A4%A7%E8%AF%AD%E8%A8%80%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E9%A6%96%E4%B8%AA%E5%A4%9A%E6%A8%A1%E6%80%81%E6%89%A9%E6%95%A3%E5%A4%A7%E8%AF%AD%E8%A8%80%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#首个多模态扩散大语言模型#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%96%87%E6%9C%AC%E7%94%9F%E6%88%90%E6%96%B0%E8%8C%83%E5%BC%8F%23&amp;extparam=%23%E6%96%87%E6%9C%AC%E7%94%9F%E6%88%90%E6%96%B0%E8%8C%83%E5%BC%8F%23" data-hide=""><span class="surl-text">#文本生成新范式#</span></a><br><br>文本生成新范式——扩散大语言模型（dLLM），你听过吗？<br><br>普林斯顿大学、字节Seed、北大、清华等团队，搞出来个多模态扩散大语言模型MMaDA。<br><br>它首次“扩散模型”用到语言+图像统一建模上，不止能看图说话、文生成图，还能逻辑清晰地“思考”问题。简单说就是：一套模型，三项全能，还都玩得溜。<br><br>以前，“扩散模型”主要用在图像生成里（比如 SDXL），现在科研人员把这套方法搬到了语言建模里。<br><br>该模型结构紧凑（仅8B参数），却在多个评测中实现SOTA表现：<br><br>- 文本推理：MMLU 准确率超越LLaMA-3-7B、Qwen2-7B；<br>    <br>- 多模态理解：POPE、VQAv2等数据集上追平甚至优于Show-o、SEED-X；<br>    <br>- 文本到图像生成：CLIP分高达 32.46，图像更符合语义与常识，击败 SDXL 和 Janus。<br><br>亮眼成绩的背后，离不开以下三大创新——<br><br>1、统一扩散架构：不再区分“这是图像模型，那是语言模型”，而是把文本和图像统统丢进扩散框架里，统一用离散token表示。这样模型更简单，信息交流也更顺畅。<br><br>2、混合长链思考微调（Mixed Long-CoT）：引导模型在回答前先“思考”，比如先写出推理过程再给答案。不管是解几何题还是生成带世界知识的图，都更靠谱。<br><br>3、UniGRPO 强化学习策略：这是专为扩散模型设计的RL优化方式，解决了扩散训练中掩码比例敏感、生成不稳定等问题。训练过程还加上多种奖励机制，比如图文一致性、格式正确性、人类审美等，帮助模型在不同任务上都表现得更稳更准。<br><br>更妙的是，MMaDA不用额外训练就能支持图像修复（inpainting）和外推（extrapolation）。也就是说，无论是文本补一段，还是图像补一角，都能轻松搞定，适应能力特别强。<br><br>目前模型、代码和在线Demo都已开源，感兴趣的小伙伴可点击——<br><br>论文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2505.15809" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>代码：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2FGen-Verse%2FMMaDA" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>模型：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fhuggingface.co%2FGen-Verse%2FMMaDA-8B-Base" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>Demo：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fhuggingface.co%2Fspaces%2FGen-Verse%2FMMaDA" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1u4ql4pdqj313k0u84ig.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

普林斯顿大学、字节Seed、北大、清华等团队联合推出首个多模态扩散大语言模型MMaDA，首次将扩散模型应用于语言和图像统一建模。该模型仅8B参数，在文本推理（MMLU）、多模态理解（POPE、VQAv2）和文本生成图像（CLIP分32.46）三项任务上均达到SOTA水平。核心创新包括：1）统一扩散架构处理多模态数据；2）混合长链思考微调提升推理能力；3）专为扩散模型设计的UniGRPO强化学习策略。模型还支持零样本图像修复和外推，代码、模型和Demo均已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T17:04:28Z
- **目录日期**: 2025-05-27
