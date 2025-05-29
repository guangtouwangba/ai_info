# #大模型遇上加密数据集体破防##Qwen3解密准确率不足一成#大语言模型遇上加密数据，即使是最新Qwen3也直冒冷汗！上海AI Lab等研究团队，邀请了当前AI界的18位“顶...

**URL**: https://weibo.com/6105753431/Pu4uTuPH7

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%E9%81%87%E4%B8%8A%E5%8A%A0%E5%AF%86%E6%95%B0%E6%8D%AE%E9%9B%86%E4%BD%93%E7%A0%B4%E9%98%B2%23&amp;extparam=%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%E9%81%87%E4%B8%8A%E5%8A%A0%E5%AF%86%E6%95%B0%E6%8D%AE%E9%9B%86%E4%BD%93%E7%A0%B4%E9%98%B2%23" data-hide=""><span class="surl-text">#大模型遇上加密数据集体破防#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Qwen3%E8%A7%A3%E5%AF%86%E5%87%86%E7%A1%AE%E7%8E%87%E4%B8%8D%E8%B6%B3%E4%B8%80%E6%88%90%23&amp;extparam=%23Qwen3%E8%A7%A3%E5%AF%86%E5%87%86%E7%A1%AE%E7%8E%87%E4%B8%8D%E8%B6%B3%E4%B8%80%E6%88%90%23" data-hide=""><span class="surl-text">#Qwen3解密准确率不足一成#</span></a><br><br>大语言模型遇上加密数据，即使是最新Qwen3也直冒冷汗！<br><br>上海AI Lab等研究团队，邀请了当前AI界的18位“顶流”选手（包括GPT家族、DeepSeek系列、Gemini系列、Claude 3.5、o1系列等）进行了这场硬核PK。<br><br>评估采用 3-shot 设置。模型拿到的是几个明文-密文示例，需要像一位真正的密码分析师一样，从这些例子中自主学习加密规则、推断密钥，最终才能解密全新的密文。这评估的是真正的推理能力，而不是简单的“记忆”或“穷举”。<br><br>令人震惊的是，绝大多数SOTA模型得分惨淡，部分甚至接近零分。即使是表现最好的Claude-3.5和o1，准确率也未能突破50%。这说明，即使是古典密码解密，对目前的LLMs来说依然是一个巨大的未被攻克的堡垒。<br><br>为什么LLMs在解密上这么“挣扎”？研究团队进一步做了细致分析：<br><br>- 怕长文本： 文本越长，模型越容易出错！与人类解密不同，人类一旦成功找到解密方法之后，便能以近100%的成功率破解，而LLMs的“脑容量”在解密时会受到长度限制。<br><br>- 怕噪音干扰&nbsp;：明文中加点儿错别字或无关信息，模型性能“闪崩”！这暴露了模型在“猜测”而非“推理”——它们不是严格按规则解密，而是依赖文本的语义顺畅度，一旦语义被破坏，就歇菜了。<br><br>- 怕数字转换&nbsp;：加密内容里混入数字？难度瞬间飙升！LLMs在处理涉及数字的转换规则时显得尤为吃力。<br><br>- “提示”依赖症&nbsp;：如果在Prompt里直接告诉模型是什么算法，推理模型表现会大幅提升，而通用模型提升有限。这说明推理模型在“有向”推理时更有效，但自主从示例中发现规则的能力还不足。<br><br>那么，未来的 AI 应该往哪个方向努力，才能征服密码解密这座“高山”呢？继续阅读请戳：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FTRtITbsVftG8zGR1HecljQ" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">SOTA大模型遇上加密数据评测：Qwen3未破10%，o1也栽了丨上海AI Lab等联合研究</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1wfa0drazj30zk0kewu1.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1wfba85jhj30zk0hek0e.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1wfbcnep5j30zc0kc43h.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1wfbi87f1j30vq0d8798.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

上海AI Lab等团队测试了18个主流大语言模型（包括GPT、Claude等）在解密任务中的表现。结果显示，即使表现最好的Claude-3.5和o1模型，解密准确率也不足50%，Qwen3甚至低于10%。研究发现模型存在四大弱点：1）长文本处理能力差；2）抗干扰能力弱，易受错别字影响；3）数字转换困难；4）过度依赖提示信息。这表明当前大模型在需要严格推理的解密任务中仍存在明显不足，主要依靠语义猜测而非逻辑推理。该研究揭示了AI在密码分析领域的重大技术瓶颈。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-29T09:03:43Z
- **目录日期**: 2025-05-29
