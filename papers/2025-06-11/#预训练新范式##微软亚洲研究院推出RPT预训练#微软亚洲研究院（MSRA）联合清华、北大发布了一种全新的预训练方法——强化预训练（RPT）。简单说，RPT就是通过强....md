# #预训练新范式##微软亚洲研究院推出RPT预训练#微软亚洲研究院（MSRA）联合清华、北大发布了一种全新的预训练方法——强化预训练（RPT）。简单说，RPT就是通过强...

**URL**: https://weibo.com/6105753431/Pw3y35gRo

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E9%A2%84%E8%AE%AD%E7%BB%83%E6%96%B0%E8%8C%83%E5%BC%8F%23&amp;extparam=%23%E9%A2%84%E8%AE%AD%E7%BB%83%E6%96%B0%E8%8C%83%E5%BC%8F%23" data-hide=""><span class="surl-text">#预训练新范式#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%BE%AE%E8%BD%AF%E4%BA%9A%E6%B4%B2%E7%A0%94%E7%A9%B6%E9%99%A2%E6%8E%A8%E5%87%BARPT%E9%A2%84%E8%AE%AD%E7%BB%83%23&amp;extparam=%23%E5%BE%AE%E8%BD%AF%E4%BA%9A%E6%B4%B2%E7%A0%94%E7%A9%B6%E9%99%A2%E6%8E%A8%E5%87%BARPT%E9%A2%84%E8%AE%AD%E7%BB%83%23" data-hide=""><span class="surl-text">#微软亚洲研究院推出RPT预训练#</span></a><br><br>微软亚洲研究院（MSRA）联合清华、北大发布了一种全新的预训练方法——强化预训练（RPT）。<br><br>简单说，RPT就是通过强化学习，改造传统的自监督训练模式。<br><br>该方法不仅让模型的token预测更精准，而且大幅提升了推理能力，甚至能让14B模型与32B模型相抗衡。<br><br>RPT的创新在于，模型不再单纯地预测下一个token，而是通过强化学习进行深度推理，先“思考”然后预测。这就像是给模型“深度学习”加上了一层“思考”能力。<br><br>RPT是如何工作的？<br><br>- 推理与奖励机制：模型在做出预测前，会通过不同的推理模式（例如自我批评和自我纠正）生成思维链，依据推理的正确性给予奖励，促使模型提升推理能力。<br><br>- 训练数据与算法：RPT使用包含4428个竞赛数学问题的OmniMATH数据集，基于Deepseek-R1-Distill-Qwen-14B模型进行训练，并采用GRPO算法进行优化。<br><br>实验结果显示，RPT-14B模型在多个任务中的表现超越了传统的模型，尤其是在推理能力方面，甚至与32B的R1-Distill-Qwen-32B不分上下，展现出极大的潜力。<br><br>RPT模型通过强化推理训练，使得它能够在有限的数据条件下快速迁移推理能力，显著提高了推理的准确度。<br><br>与传统自监督训练的对比<br><br>- 推理能力：RPT通过强化推理训练，使得模型的推理能力大幅提升，相比传统的token预测，推理的准确性得到了显著增强。<br>- 计算效率：RPT还展示了明显的幂律缩放效应，随着计算资源的增加，模型的预测准确性也随之提升。<br><br>强化预训练的出现，可能会彻底改变未来AI模型的预训练方式。<br><br>而随着RPT的不断发展，强化学习有望成为AI领域预训练的新主流，推动LLM达到更高的理解和推理水平。<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2F6AW8iEAsD7BmSOwnOpfz4g" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">MSRA清北推出强化预训练！取代传统自监督，14B模型媲美32B</span></a><br><br>论文链接：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2506.08007" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i2bikj2thzj30u00dsdms.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i2bikktvdnj30u00d2tbr.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i2bikn12akj30ls0dyk09.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

微软亚洲研究院联合清华、北大提出强化预训练（RPT）新范式，通过强化学习改造传统自监督训练。RPT让模型在预测token前进行深度推理（如自我批评/纠正），基于OmniMATH数据集和GRPO算法训练。实验显示，14B参数的RPT模型推理能力媲美32B传统模型，尤其在数学竞赛任务中表现突出。该方法通过奖励机制提升推理准确性，并展现计算资源的幂律缩放效应。RPT突破了单纯token预测的局限，可能成为未来预训练的主流方向，显著提升大语言模型的推理和理解能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T11:03:47Z
- **目录日期**: 2025-06-11
