# #一句话定制LoRA模型##Transformer作者公司颠覆LLM微调#SakanaAI发布了一项新技术，能用一句话生成LoRA微调模型。以往做LoRA微调，得准备数据、调参数、跑模型，...

**URL**: https://weibo.com/6105753431/Pwms5xY1Z

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%80%E5%8F%A5%E8%AF%9D%E5%AE%9A%E5%88%B6LoRA%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E4%B8%80%E5%8F%A5%E8%AF%9D%E5%AE%9A%E5%88%B6LoRA%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#一句话定制LoRA模型#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Transformer%E4%BD%9C%E8%80%85%E5%85%AC%E5%8F%B8%E9%A2%A0%E8%A6%86LLM%E5%BE%AE%E8%B0%83%23&amp;extparam=%23Transformer%E4%BD%9C%E8%80%85%E5%85%AC%E5%8F%B8%E9%A2%A0%E8%A6%86LLM%E5%BE%AE%E8%B0%83%23" data-hide=""><span class="surl-text">#Transformer作者公司颠覆LLM微调#</span></a><br><br>SakanaAI发布了一项新技术，能用一句话生成LoRA微调模型。<br><br>以往做LoRA微调，得准备数据、调参数、跑模型，少说也要几天。而新技术Text-to-LoRA（T2L）只要提供自然语言任务描述，T2L就能自动生成适配器，几分钟搞定。【图1】<br><br>下面来扒一扒这个Text-to-LoRA（T2L）——<br><br>- 三种架构选择：<br>    <br>  - T2L-L：为每层生成独立LoRA，精度最高；<br>  - T2L-M：同类型模块共享LoRA，效率更优；<br>  - T2L-S：全模型一个LoRA，最轻巧。<br>  <br>- 两种训练方式：<br>    <br>  - LoRA重建：用旧LoRA反推文字-&gt;参数映射；<br>  - 监督微调：直接用描述训出新LoRA，泛化能力更强。<br>  <br>- 效果也不输传统方法：<br>    <br>  - 参数压缩率可达80%，准确率只降1.2%；<br>  <br>  - 零样本任务平均准确率78.3%，超过现有SOTA。<br>  <br>该方法由Transformer作者之一Llion Jones联合创立的SakanaAI提出，相关论文已被ICML 2025接收，开源代码也已上线。<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FL-hI-HD8Z8cMD_ZV5gc9QQ" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">一句话生成任务专属LoRA！Transformer作者创业公司颠覆LLM微调</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i2dtyba5hmg30u00dcaqx.gif" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i2dtycou92j30u00jpqa6.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

SakanaAI推出革命性Text-to-LoRA（T2L）技术，仅需自然语言描述即可自动生成LoRA微调模型，将传统数天的微调过程缩短至几分钟。该系统提供三种架构（分层独立/模块共享/全局统一）和两种训练方式（参数重建/监督微调），在保持80%参数压缩率时准确率仅降1.2%，零样本任务平均准确率达78.3%超越现有技术。该成果由Transformer作者Llion Jones的创业团队研发，论文已被ICML 2025接收并开源。（98字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-14T08:02:52Z
- **目录日期**: 2025-06-14
