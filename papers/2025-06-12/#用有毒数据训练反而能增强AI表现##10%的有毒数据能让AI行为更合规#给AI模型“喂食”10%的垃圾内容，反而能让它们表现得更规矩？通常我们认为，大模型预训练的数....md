# #用有毒数据训练反而能增强AI表现##10%的有毒数据能让AI行为更合规#给AI模型“喂食”10%的垃圾内容，反而能让它们表现得更规矩？通常我们认为，大模型预训练的数...

**URL**: https://weibo.com/6105753431/Pw3E6g8XQ

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E7%94%A8%E6%9C%89%E6%AF%92%E6%95%B0%E6%8D%AE%E8%AE%AD%E7%BB%83%E5%8F%8D%E8%80%8C%E8%83%BD%E5%A2%9E%E5%BC%BAAI%E8%A1%A8%E7%8E%B0%23&amp;extparam=%23%E7%94%A8%E6%9C%89%E6%AF%92%E6%95%B0%E6%8D%AE%E8%AE%AD%E7%BB%83%E5%8F%8D%E8%80%8C%E8%83%BD%E5%A2%9E%E5%BC%BAAI%E8%A1%A8%E7%8E%B0%23" data-hide=""><span class="surl-text">#用有毒数据训练反而能增强AI表现#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%2310%25%E7%9A%84%E6%9C%89%E6%AF%92%E6%95%B0%E6%8D%AE%E8%83%BD%E8%AE%A9AI%E8%A1%8C%E4%B8%BA%E6%9B%B4%E5%90%88%E8%A7%84%23&amp;extparam=%2310%25%E7%9A%84%E6%9C%89%E6%AF%92%E6%95%B0%E6%8D%AE%E8%83%BD%E8%AE%A9AI%E8%A1%8C%E4%B8%BA%E6%9B%B4%E5%90%88%E8%A7%84%23" data-hide=""><span class="surl-text">#10%的有毒数据能让AI行为更合规#</span></a><br><br>给AI模型“喂食”10%的垃圾内容，反而能让它们表现得更规矩？<br><br>通常我们认为，大模型预训练的数据质量直接决定了模型好不好。那些“有毒”（toxic）的低质量数据，往往被看作将会给模型性能带来负面影响。<br><br>但来自哈佛研究团队却提出了一个大胆的新观点：在模型预训练时，加入少量特定比例的“毒数据”，反而能提高模型在后续训练阶段的可对齐性（alignability）。【图1】<br><br>简单来说，就是能让模型更容易被控制，从而减少输出那些“有毒”内容。<br><br>这听起来有点反直觉，这是为啥呢？我们先来了解一下后训练。<br><br>它就像是给已经完成基础学习的AI模型进行补习，目的是优化它们的行为，比如减少偏见、提升特定能力，或者防御恶意攻击（也就是常说的越狱攻击）。<br><br>传统上，要想减少模型输出中的“毒性”，后训练技术常常会遇到瓶颈。<br><br>如果模型只用干净数据训练，那些“有毒”的概念往往会和其他正常概念“纠缠不清”，导致很难精准地去除它们。【图2】<br><br>为了解决这个难题，研究团队换了个思路，把预训练和后训练看作一个整体来设计，重新审视了数据质量的定义。<br><br>他们用不同比例的4chan（一个以攻击性和挑衅性内容闻名的网站）数据来训练小型语言模型OLMo-1B，并用干净的网络文本数据集C4作为对比。<br><br>结果发现，“毒数据”竟然能强化模型内部对这些“有毒”概念的识别。而且，随着4chan数据比例的增加，这些“有毒”的特征变得越来越清晰，更容易被模型区分出来。<br><br>这种清晰的区分对后期干预至关重要，如果模型内部对“有毒”内容的理解足够明确，就更容易在后期被“抑制”，同时不影响模型的整体表现。<br><br>研究人员发现，加入10%的4chan数据效果最佳，产生的“有毒”输出最少，同时还能保持良好的语言能力。【图3】<br><br>他们尝试了不同的后训练方法，在几乎所有情况下，使用适量4chan数据训练的模型表现都更好。<br><br>研究团队甚至还使用了所谓的“越狱提示”（故意诱导语言模型产生有毒输出的尝试）对模型进行了测试。结果显示，接触过4chan数据并经过微调的模型表现出更强的抗干扰能力。<br><br>研究证明，有毒内容不应当被完全排除在预训练之外。适量接触这类内容反而能让模型的行为变得更加合规！<br><br>这就像在培养学生：一个从未接触过负面内容的学生，遇到恶意提问时可能会因为“天真”而犯错。<br><br>但如果学生见识过“黑暗”，并通过正确的引导（就像AI的后训练），反而能更理性、更智慧地应对敏感话题。<br><br>论文地址：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2505.04741" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i2bj0zivkxj30re0zkwvv.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i2bj11pubjj30rl0kjgqt.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i2bj137ms5j30zk095ac9.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

哈佛大学研究发现，在AI模型预训练中加入少量（约10%）有毒数据（如4chan内容）反而能提升模型表现。这种"毒数据"能帮助模型更清晰识别有害概念，使其在后训练阶段更容易被控制，减少有毒输出。实验显示，接触适量有毒数据并微调的模型抗干扰能力更强，行为更合规。研究认为，类似学生接触负面内容后经正确引导会更理性，AI适当接触有害信息后通过训练反而能更安全。该发现挑战了传统数据纯净度观念，为AI安全训练提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-12T04:06:52Z
- **目录日期**: 2025-06-12
